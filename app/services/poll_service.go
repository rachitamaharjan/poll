package services

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/models"
)

var jobQueue = make(chan models.VoteJob, 100) // Initialize the job queue once
var once sync.Once

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
		_, err := models.UpdatePollVotes(job.PollID, job.OptionIndex)
		if err != nil {
			log.Printf("Failed to update vote: %v", err)
			continue
		}

	}
}
