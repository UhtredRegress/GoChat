package handler

import (
	"github.com/UhtredRegress/GoChat/internal/dto"
	"github.com/gin-gonic/gin"
)

func createUserHandler(c *gin.Context) {
	var createUserRequest dto.CreateUserRequest
	if err := c.BindJSON(&createUserRequest); err != nil {
		c.JSON(400, "The request is not valid")
		return
	}
	c.JSON(200, "Success")
}



func RegisterUserHandler(e *gin.Engine) {
	userApi := e.Group("/api/v1/user")
	{
		userApi.POST("create", createUserHandler)
	}
}
