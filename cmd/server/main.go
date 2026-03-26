package main

import (
	"github.com/gin-gonic/gin"
	"github.com/UhtredRegress/GoChat/internal/handler"
)

func main() {
	r := gin.Default()
	handler.RegisterUserHandler(r)
}
