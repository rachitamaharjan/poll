package services

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rachitamaharjan/poll/db"
	"github.com/rachitamaharjan/poll/models"
	"github.com/sirupsen/logrus"
)

var jobQueue = make(chan models.VoteJob, 100) // Initialize the job queue once
var once sync.Once
var pollSubscribers = make(map[uuid.UUID][]chan string) // Keyed by Poll ID

func init() {
	once.Do(func() {
		go processJobQueue() // Start the goroutine to process the queue
	})
}

func GetPollByID(c *gin.Context, pollId uuid.UUID) (*models.Poll, error) {
	poll, err := models.GetPollByID(c, pollId)
	if err != nil {
		return nil, err
	}
	return poll, nil
}

func CreatePoll(c *gin.Context, newPoll models.Poll) (string, error) {
	newPoll.ID = uuid.New()

	// Save to db
	_, err := models.SavePoll(c, newPoll)
	if err != nil {
		return "", err
	}
	return newPoll.ID.String(), nil
}

func VotePoll(c *gin.Context, pollId uuid.UUID, vote models.VoteRequest) *models.CustomError {
	// Enqueue the vote job
	err := enqueueVoteJob(c, pollId, vote.OptionIndex)
	if err != nil {
		return err
	}
	return nil
}

func enqueueVoteJob(c *gin.Context, pollId uuid.UUID, optionIndex int) *models.CustomError {
	clientIP := c.ClientIP()
	// Check for IP-based voting restrictions
	poll, err := models.GetPollByID(c, pollId)
	if err != nil {
		return &models.CustomError{Message: err.Error()}
	}
	if !poll.AllowMultipleVotes {
		var existingVote models.PollVote
		err = db.DB.Where("poll_id = ? AND ip_address = ?", pollId, clientIP).First(&existingVote).Error
		if err == nil {
			return &models.CustomError{Message: models.ErrMultipleVotesNotAllowed, StatusCode: 400}
		}
	}

	// Enqueue the vote job
	jobQueue <- models.VoteJob{
		PollID:      pollId,
		OptionIndex: optionIndex,
		ClientIP:    clientIP,
	}
	return nil
}

func processJobQueue() {
	for job := range jobQueue {
		poll, jobErr := models.UpdatePollVotes(job.PollID, job.OptionIndex, job.ClientIP)
		if jobErr != nil {
			logrus.Errorf("Failed to update vote: %v", jobErr)
			continue
		}

		// Notify all subscribers of this poll (SSE clients)
		pollJSON, err := json.Marshal(poll)
		if err == nil {
			fmt.Print("pollSubscribers ", pollSubscribers, "pollid", poll.ID)
			if subscribers, ok := pollSubscribers[poll.ID]; ok {
				for _, subscriber := range subscribers {
					select {
					case subscriber <- string(pollJSON):
						logrus.Infof("Sent update to subscriber for poll %v", poll.ID)
					default:
						logrus.Warn("Failed to send update to subscriber")
					}
				}
			} else {
				logrus.Warn("No subscribers found for poll %v", poll.ID)
			}
		} else {
			logrus.Errorf("Failed to marshal poll JSON: %v", err)
		}
	}
}

func PollsStream(c *gin.Context, pollID string) {
	// Create a new channel for this client
	updateChannel := make(chan string)

	// Add the new subscriber to the pollSubscribers map
	pollUUID, err := uuid.Parse(pollID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid poll ID"})
		return
	}

	pollSubscribers[pollUUID] = append(pollSubscribers[pollUUID], updateChannel)

	// Remove the client when they disconnect
	defer func() {
		subscribers := pollSubscribers[pollUUID]
		for i, subscriber := range subscribers {
			if subscriber == updateChannel {
				pollSubscribers[pollUUID] = append(subscribers[:i], subscribers[i+1:]...)
				break
			}
		}
		close(updateChannel)
	}()

	// Listen for poll updates and send them to the client
	for {
		select {
		case pollUpdate := <-updateChannel:
			_, err := fmt.Fprintf(c.Writer, "data: %s\n\n", pollUpdate)
			if err != nil {
				logrus.Info("Client disconnected")
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			logrus.Info("Client disconnected")
			return
		}
	}
}
