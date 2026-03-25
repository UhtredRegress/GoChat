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
	api := e.Group("/api")
	v1 := api.Group("/v1")
	userApi := v1.Group("/user")
	{
		userApi.POST("create", createUserHandler)
	}
}
