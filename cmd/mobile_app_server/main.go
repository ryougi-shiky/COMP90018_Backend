package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryougi-shiky/COMP90018_Backend/api/routesHandlers"
	"github.com/ryougi-shiky/COMP90018_Backend/repository"
	"github.com/ryougi-shiky/COMP90018_Backend/services"
)

func main() {
	fmt.Println("Starting server...")
	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)
	router := gin.Default()

	router.POST("/user/register", routesHandlers.RegisterUserHandler(userService))
	router.GET("user/login", routesHandlers.LoginUserHandler(userService))

	log.Fatal(http.ListenAndServe(":8080", router))
}
