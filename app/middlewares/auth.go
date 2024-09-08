package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rachitamaharjan/poll/db"
	"github.com/rachitamaharjan/poll/models"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		var errMsg string
		switch {
		case errors.Is(err, jwt.ErrTokenMalformed):
			errMsg = "Invalid token"
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			errMsg = "Invalid signature"
		case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
			errMsg = "Token expired or not active yet"
		default:
			errMsg = fmt.Sprintf("Error handling token: %v", err)
		}
		fmt.Println(errMsg)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Find the user with token sub
		var user models.User
		if sub, ok := claims["sub"].(float64); ok {
			db.DB.Where(&models.User{ID: uint(sub)}).First(&user)
			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			// Attach user to request
			c.Set("user", user)
			c.Next()
			return
		}
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}
