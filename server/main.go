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
	db, err := repository.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	memoRepository := repository.NewMemoRepository(db)

	userService := services.NewUserService(userRepository)
	memoService := services.NewMemoService(memoRepository)

	router := gin.Default()

	router.POST("/user/register", routesHandlers.RegisterUserHandler(userService))
	router.GET("/user/login", routesHandlers.LoginUserHandler(userService))

	router.POST("/memo/create", routesHandlers.CreateMemo(memoService))
	router.PUT("/memo/update", routesHandlers.UpdateMemo(memoService))
	router.DELETE("/memo/delete", routesHandlers.DeleteMemo(memoService))
	router.GET("/memo/getmemos", routesHandlers.GetUserMemo(memoService))

	log.Fatal(http.ListenAndServe(":8080", router))
}
