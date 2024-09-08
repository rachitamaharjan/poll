package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/constants"
	handlers "github.com/rachitamaharjan/poll/handlers"
	"github.com/rachitamaharjan/poll/middlewares"
)

func AuthController(router *gin.RouterGroup) {
	router.POST(constants.SIGN_UP_ROUTE, handlers.SignUp)
	router.POST(constants.LOGIN_ROUTE, handlers.LogIn)
	router.GET("/validate", middlewares.RequireAuth, handlers.Validate)
}
