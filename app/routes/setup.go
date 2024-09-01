package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/constants"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group(constants.API_V1)

	setupAuthRoutes(v1)
	setupPollRoutes(v1)
	return router
}
