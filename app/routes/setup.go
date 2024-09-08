package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/constants"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group(constants.API_V1)

	setupAuthRoutes(v1)
	setupPollRoutes(v1)
	return router
}
