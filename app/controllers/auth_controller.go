package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/leave-management-system/constants"
	handlers "github.com/rachitamaharjan/leave-management-system/handlers"
	"github.com/rachitamaharjan/leave-management-system/middlewares"
)

func AuthController(router *gin.RouterGroup) {
	router.POST(constants.SIGN_UP_ROUTE, handlers.SignUp)
	router.POST(constants.LOGIN_ROUTE, handlers.LogIn)
	router.GET("/validate", middlewares.RequireAuth, handlers.Validate)
}
