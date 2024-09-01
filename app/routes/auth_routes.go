package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/constants"
	"github.com/rachitamaharjan/leave-management-system/controllers"
)

func setupAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group(constants.AUTH_GROUP)
	{
		controllers.AuthController(authGroup)
	}
}
