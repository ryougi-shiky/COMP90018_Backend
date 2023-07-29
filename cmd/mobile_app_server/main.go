package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"github.com/ryougi-shiky/COMP90018_Backend/repository"
	"github.com/ryougi-shiky/COMP90018_Backend/services"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	fmt.Println("Starting server...")
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)
	router := gin.Default()

	router.POST("/user/register", func (c *gin.Context)  {
		// Register a new user
		var request RegisterRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := models.User {
			ID: uuid.New(),
			Username: request.Username,
			Email: request.Email,
			Password: request.Password,
		}

		err := userService.RegisterUser(&user)
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}