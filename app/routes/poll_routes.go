package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/constants"
	"github.com/rachitamaharjan/leave-management-system/controllers"
)

func setupPollRoutes(router *gin.RouterGroup) {
	pollGroup := router.Group(constants.POLLS_GROUP)
	{
		controllers.PollController(pollGroup)
	}
}
