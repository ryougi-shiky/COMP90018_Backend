package web

import (
	"github.com/gin-gonic/gin"
	"github.com/ryougi-shiky/COMP90018_Backend/services"
)

func RegisterUserHandler(userService services.UserService) gin.HandlerFunc{
	return func (c *gin.Context)  {
		// Parse user from request and call userService.RegisterUser
	}
}