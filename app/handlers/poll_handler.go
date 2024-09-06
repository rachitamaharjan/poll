package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/models"
)
var (
	jobQueue chan models.VoteJob
)

// Get poll by ID
func GetPollByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	poll, err := models.GetPollByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, poll)
}

// Create a new poll
func CreatePoll(c *gin.Context) {
	var newPoll models.Poll
	if err := c.ShouldBindJSON(&newPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save to db
	pollID, err := models.SavePoll(c, newPoll)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Return a unique poll URL
	c.JSON(http.StatusCreated, gin.H{"url": fmt.Sprintf("/polls/%d", pollID)})
}

// VotePoll handles the voting request from a user.
func VotePoll(c *gin.Context) {
	pollId, _ := strconv.Atoi(c.Param("id"))
	var vote struct {
		OptionIndex int `json:"optionIndex"`
	}

	// TODO: handle error if no option selected
	// TODO: also validate JSON data and ensure only one option is selected
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	c.JSON(http.StatusOK, gin.H{"message": "Vote recorded"})
}
