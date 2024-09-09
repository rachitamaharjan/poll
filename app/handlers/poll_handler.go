package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/models"
	"github.com/rachitamaharjan/poll/services"
)

// Get poll by ID
func GetPollByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	poll, err := services.GetPollByID(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	pollID, err := services.CreatePoll(c, newPoll)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Return a unique poll URL
	c.JSON(http.StatusCreated, gin.H{"url": fmt.Sprintf("/polls/%d", pollID)})
}

// VotePoll handles the voting request from a user.
func VotePoll(c *gin.Context) {
	pollId, _ := strconv.Atoi(c.Param("id"))
	var vote models.VoteRequest

	// TODO: handle error if no option selected
	// TODO: also validate JSON data and ensure only one option is selected
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.VotePoll(pollId, vote)

	c.JSON(http.StatusOK, gin.H{"message": "Vote recorded"})
}

func PollsStream(c *gin.Context) {
	pollID := c.Param("id")

	// Set the headers to indicate an SSE stream
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	services.PollsStream(c, pollID)
}
