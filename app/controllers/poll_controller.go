package controllers

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/rachitamaharjan/leave-management-system/handlers"
)

func PollController(router *gin.RouterGroup) {
	router.GET("/", handlers.GetPolls)
	router.POST("/", handlers.CreatePoll)
	router.GET("/:id", handlers.GetPollByID)
	router.POST("/:id/vote", handlers.VotePoll)
}
