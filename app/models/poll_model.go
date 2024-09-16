package models

import (
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rachitamaharjan/poll/db"
	"github.com/sirupsen/logrus"
)

var (
	mu sync.Mutex // Mutex to ensure thread-safe operations
)

var (
	ErrMultipleVotesNotAllowed = "Only one vote allowed per IP"
)

// Poll represents a poll with a question and multiple options.
type Poll struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	Question           string   `json:"question" gorm:"type:text"`
	Options            []Option `json:"options" gorm:"foreignKey:PollID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedBy          int      `json:"createdBy"`
	AllowMultipleVotes bool     `json:"allowMultipleVotes"`
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
	ClientIP    string
}

// PollVote represents a record of a vote by an IP address.
type PollVote struct {
	ID        uint      `gorm:"primaryKey"`
	PollID    uuid.UUID `json:"pollId"`
	IPAddress string    `json:"ipAddress" gorm:"type:varchar(45)"` // Supports both IPv4 and IPv6
	CreatedAt time.Time
}

type CustomError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func GetPollByID(c *gin.Context, id uuid.UUID) (*Poll, error) {
	poll := &Poll{}
	result := db.DB.Preload("Options").First(poll, id)
	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"poll_id": id,
			"error":   result.Error,
		}).Error("Failed to fetch poll")
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
		logrus.WithFields(logrus.Fields{
			"poll":  poll.Question,
			"error": result.Error,
		}).Error("Failed to create poll")
		return "", result.Error
	}

	logrus.WithFields(logrus.Fields{
		"poll_id": poll.ID,
		"poll":    poll.Question,
	}).Info("Poll created successfully")

	return poll.ID.String(), nil
}

// UpdatePollVotes increments the vote count for a specific option in the poll.
func UpdatePollVotes(pollID uuid.UUID, optionIndex int, clientIP string) (*Poll, *CustomError) {
	var option Option
	var poll Poll

	// Fetch the option to be updated
	err := db.DB.Where("poll_id = ?", pollID).
		Offset(optionIndex). // Skip to the correct index
		Limit(1).            // Get only one option
		First(&option).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"poll_id":      pollID,
			"option_index": optionIndex,
			"error":        err,
		}).Error("Failed to fetch option")
		return nil, &CustomError{Message: err.Error()}
	}

	// For thread safety using mutex
	mu.Lock()
	defer mu.Unlock()

	// Increment the vote count for the selected option
	option.VoteCount += 1

	// Save the updated option
	err = db.DB.Save(&option).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"option_id": option.ID,
			"poll_id":   pollID,
			"error":     err,
		}).Error("Failed to update option votes")
		return nil, &CustomError{Message: err.Error()}
	}

	// Fetch the updated poll with its options
	poll, err = fetchUpdatedPollWithOptions(poll, pollID)
	if err != nil {
		return nil, &CustomError{Message: err.Error()}
	}

	logrus.WithFields(logrus.Fields{
		"poll_id": pollID,
	}).Info("Poll votes updated successfully")

	recordVoteIfNecessary(poll, pollID, clientIP)

	return &poll, nil
}

func multipleVotesAllowed(poll Poll, pollID uuid.UUID) (bool, error) {
	// Fetch the poll to check if multiple votes are allowed
	err := db.DB.Where("id = ?", pollID).Preload("Options").First(&poll).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"poll_id": pollID,
			"error":   err,
		}).Error("Failed to fetch poll")
		return false, err
	}

	return poll.AllowMultipleVotes, nil
}

func fetchUpdatedPollWithOptions(poll Poll, pollID uuid.UUID) (Poll, error) {
	err := db.DB.Where("id = ?", pollID).Preload("Options").First(&poll).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"poll_id": pollID,
			"error":   err,
		}).Error("Failed to fetch updated poll")
		return Poll{}, err
	}
	// Sort the options by VoteCount in descending order
	sort.Slice(poll.Options, func(i, j int) bool {
		return poll.Options[i].VoteCount > poll.Options[j].VoteCount
	})
	return poll, nil
}

func recordVoteIfNecessary(poll Poll, pollID uuid.UUID, clientIP string) {
	if !poll.AllowMultipleVotes {
		db.DB.Create(&PollVote{
			PollID:    pollID,
			IPAddress: clientIP,
		})
	}
}
