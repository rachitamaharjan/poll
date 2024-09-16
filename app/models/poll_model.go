package models

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rachitamaharjan/poll/db"
)

var (
	mu sync.Mutex // Mutex to ensure thread-safe operations
)

// Poll represents a poll with a question and multiple options.
type Poll struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Question  string   `json:"question" gorm:"type:text"`
	Options   []Option `json:"options" gorm:"foreignKey:PollID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy int      `json:"created_by"`
}

// Option represents an individual poll option with a vote count.
type Option struct {
	ID        uint      `gorm:"primaryKey"`
	PollID    uuid.UUID `json:"pollId"`
	Text      string    `json:"text" gorm:"type:text"`
	VoteCount int       `json:"voteCount" gorm:"default:0"`
}

type VoteRequest struct {
	OptionIndex int `json:"optionIndex"`
}

type VoteJob struct {
	PollID      uuid.UUID
	OptionIndex int
}

func GetPollByID(c *gin.Context, id uuid.UUID) (*Poll, error) {
	poll := &Poll{}
	result := db.DB.Preload("Options").First(poll, id)
	if result.Error != nil {
		log.Printf("Failed to create poll. Error: %v", result.Error)
		return nil, result.Error
	}
	// Sort the options by VoteCount in descending order
	sort.Slice(poll.Options, func(i, j int) bool {
		return poll.Options[i].VoteCount > poll.Options[j].VoteCount
	})
	return poll, nil
}

// Saves a poll
func SavePoll(c *gin.Context, poll Poll) (string, error) {
	result := db.DB.Create(&poll)
	if result.Error != nil {
		log.Printf("Failed to create poll. Error: %v", result.Error)
		return "", result.Error
	}
	return poll.ID.String(), nil
}

// UpdatePollVotes increments the vote count for a specific option in the poll.
func UpdatePollVotes(pollID uuid.UUID, optionIndex int) (*Poll, error) {
	var option Option
	var poll Poll

	// Fetch the option to be updated
	err := db.DB.Where("poll_id = ?", pollID).
		Offset(optionIndex). // Skip to the correct index
		Limit(1).            // Get only one option
		First(&option).Error
	if err != nil {
		return nil, err
	}

	// For thread safety using mutex
	mu.Lock()
	defer mu.Unlock()

	// Increment the vote count for the selected option
	option.VoteCount += 1

	// Save the updated option
	err = db.DB.Save(&option).Error
	if err != nil {
		return nil, err
	}

	// Fetch the updated poll with its options
	err = db.DB.Where("id = ?", pollID).Preload("Options").First(&poll).Error
	if err != nil {
		return nil, err
	}

	return &poll, nil
}
