package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/models"
)

var jobQueue = make(chan models.VoteJob, 100) // Initialize the job queue once
var once sync.Once
var pollSubscribers = make(map[uint][]chan string) // Keyed by Poll ID

func init() {
	once.Do(func() {
		go processJobQueue() // Start the goroutine to process the queue
	})
}

func GetPollByID(c *gin.Context, pollId int) (*models.Poll, error) {
	poll, err := models.GetPollByID(c, uint(pollId))
	if err != nil {
		return nil, err
	}
	return poll, nil
}

func CreatePoll(c *gin.Context, newPoll models.Poll) (int, error) {
	// Save to db
	pollID, err := models.SavePoll(c, newPoll)
	if err != nil {
		return 0, err
	}
	return pollID, nil
}

func VotePoll(pollId int, vote models.VoteRequest) {
	// Enqueue the vote job
	jobQueue <- models.VoteJob{
		PollID:      uint(pollId),
		OptionIndex: vote.OptionIndex,
	}
}

func processJobQueue() {
	for job := range jobQueue {
		poll, err := models.UpdatePollVotes(job.PollID, job.OptionIndex)
		if err != nil {
			log.Printf("Failed to update vote: %v", err)
			continue
		}

		// Notify all subscribers of this poll (SSE clients)
		pollJSON, err := json.Marshal(poll)
		if err == nil {
			if subscribers, ok := pollSubscribers[poll.ID]; ok {
				for _, subscriber := range subscribers {
					select {
					case subscriber <- string(pollJSON):
					default:
						log.Printf("Failed to send update to subscriber")
					}
				}
			}
		} else {
			log.Printf("Failed to marshal poll JSON: %v", err)
		}
	}
}

func PollsStream(c *gin.Context, pollID string) {
	// Create a new channel for this client
	updateChannel := make(chan string)

	// Add the new subscriber to the pollSubscribers map
	pollIDUint, _ := strconv.ParseUint(pollID, 10, 32)
	pollSubscribers[uint(pollIDUint)] = append(pollSubscribers[uint(pollIDUint)], updateChannel)

	// Remove the client when they disconnect
	defer func() {
		subscribers := pollSubscribers[uint(pollIDUint)]
		for i, subscriber := range subscribers {
			if subscriber == updateChannel {
				pollSubscribers[uint(pollIDUint)] = append(subscribers[:i], subscribers[i+1:]...)
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
				log.Printf("Client disconnected")
				return
			}
			c.Writer.Flush()
		case <-c.Request.Context().Done():
			log.Printf("Client disconnected")
			return
		}
	}
}
