package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/constants"
	"github.com/rachitamaharjan/poll/controllers"
)

func setupPollRoutes(router *gin.RouterGroup) {
	pollGroup := router.Group(constants.POLLS_GROUP)
	{
		controllers.PollController(pollGroup)
	}
}
