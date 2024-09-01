package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/models"
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

// Vote on a poll
func VotePoll(c *gin.Context) {
	pollId, _ := strconv.Atoi(c.Param("id"))
	var vote struct {
		OptionIndex int `json:"optionIndex"`
	}
	// TODO: handle error if no option selected
	// TODO: also validate json data
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Increment the vote count for the selected option
	models.UpdatePollVotes(uint(pollId), vote.OptionIndex)
	c.JSON(http.StatusOK, gin.H{"message": "Vote recorded"})
}
