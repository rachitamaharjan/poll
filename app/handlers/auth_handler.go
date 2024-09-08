package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rachitamaharjan/poll/models"
	"github.com/rachitamaharjan/poll/services"
)

func SignUp(c *gin.Context) {
	signUpRequest := models.SignUpRequest{}
	err := c.Bind(&signUpRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})
		return
	}

	err = services.SignUp(c, signUpRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to SignUp" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func LogIn(c *gin.Context) {
	loginRequest := models.LoginRequest{}
	err := c.Bind(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read Body",
		})
		return
	}

	tokenString, err := services.LogIn(c, loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to login" + err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 2600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in",
	})
}
