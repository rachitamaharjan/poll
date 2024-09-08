package services

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/models"
)

var (
	jobQueue chan models.VoteJob
)

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

	// Initialize the job queue
	jobQueue = make(chan models.VoteJob, 100) // Buffered channel to handle 100 votes at a time

	// Enqueue the vote job
	jobQueue <- models.VoteJob{
		PollID:      uint(pollId),
		OptionIndex: vote.OptionIndex,
	}

	// Increment the vote count for the selected option
	go func() {
		for job := range jobQueue {
			if err := models.UpdatePollVotes(job.PollID, job.OptionIndex); err != nil {
				log.Printf("Failed to update vote: %v", err)
			}
		}
	}()
}
