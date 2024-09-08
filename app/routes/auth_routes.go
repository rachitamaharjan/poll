package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/constants"
	"github.com/rachitamaharjan/poll/controllers"
)

func setupAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group(constants.AUTH_GROUP)
	{
		controllers.AuthController(authGroup)
	}
}
