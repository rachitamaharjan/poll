package models

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/db"
)

type SignUpRequest struct {
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

// Creates a user
func CreateUser(r SignUpRequest, hash []byte) User {
	return User{
		Email:    r.Email,
		Password: string(hash),
	}
}

// Saves a user
func SaveUser(c *gin.Context, r SignUpRequest, user User) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Failed to create user. Error: %v", result.Error)
		return result.Error
	}
	return nil
}
