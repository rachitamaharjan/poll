package models

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/db"
	"gorm.io/gorm"
)

// Poll represents a poll with a question and multiple options.
type Poll struct {
	gorm.Model
	ID        uint `json:"id" gorm:"primaryKey;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Question  string   `json:"question" gorm:"type:text"`
	Options   []Option `json:"options" gorm:"foreignKey:PollID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy int      `json:"created_by"`
}

// Option represents an individual poll option with a vote count.
type Option struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	PollID    uint
	Text      string `gorm:"type:text"`
	VoteCount int    `json:"voteCount" gorm:"default:0"`
}

func GetPollByID(c *gin.Context, id uint) (*Poll, error) {
	poll := &Poll{}
	result := db.DB.Preload("Options").First(poll, id)
	if result.Error != nil {
		log.Printf("Failed to create poll. Error: %v", result.Error)
		return nil, result.Error
	}
	return poll, nil
}

// Saves a poll
func SavePoll(c *gin.Context, poll Poll) (int, error) {
	result := db.DB.Create(&poll)
	if result.Error != nil {
		log.Printf("Failed to create poll. Error: %v", result.Error)
		return 0, result.Error
	}
	return int(poll.ID), nil
}

// UpdatePollVotes increments the vote count for a specific option in the poll.
func UpdatePollVotes(pollID uint, optionIndex int) error {
	var option Option

	err := db.DB.Where("poll_id = ?", pollID).
		Offset(optionIndex). // Skip to the correct index
		Limit(1).            // Get only one option
		First(&option).Error
	if err != nil {
		return err
	}

	option.VoteCount += 1

	return db.DB.Save(&option).Error
}
