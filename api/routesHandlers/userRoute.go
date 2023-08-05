package routesHandlers

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"github.com/ryougi-shiky/COMP90018_Backend/services"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUserHandler(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RegisterRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User{
			ID:       uuid.New(),
			Username: request.Username,
			Email:    request.Email,
			Password: request.Password,
		}

		err := userService.RegisterUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	}
}

func LoginUserHandler(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request LoginRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.GetUserByEmail(request.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		hash := sha256.New()
		hash.Write([]byte(request.Password))
		hashedPassword := hex.EncodeToString(hash.Sum(nil))

		if user.Password != hashedPassword {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
	}
}
