package controllers

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/rachitamaharjan/poll/handlers"
)

func PollController(router *gin.RouterGroup) {
	router.POST("/", handlers.CreatePoll)
	router.GET("/:id", handlers.GetPollByID)
	router.POST("/:id/vote", handlers.VotePoll)

	// SSE endpoint for poll updates
	router.GET("/:id/stream", handlers.PollsStream)
}
