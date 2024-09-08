package services

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/rachitamaharjan/poll/db"
	"github.com/rachitamaharjan/poll/models"
)

func SignUp(c *gin.Context, r models.SignUpRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	if err != nil {
		log.Printf("Failed to hash password. Error: %v", err)
		return err
	}

	user := models.CreateUser(r, hash)

	return models.SaveUser(c, r, user)
}

func LogIn(c *gin.Context, r models.LoginRequest) (string, error) {
	user := models.User{}
	db.DB.Where(&models.User{Email: r.Email}).First(&user)
	if user.ID == 0 {
		log.Print("Invalid email or password")
		return "", nil
	}

	// user auth
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		log.Printf("Failed to get SQL database connection. Error: %v", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("Failed to create token Error: %v", err)
		return "", err
	}
	return tokenString, nil
}
