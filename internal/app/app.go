package app

import (
	"github.com/UhtredRegress/GoChat/internal/handler"
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine	
}

func NewApp() *App {
	router := gin.Default()
	return &App{router}
}

func (a *App) RegisterHandler() {
	handler.RegisterUserHandler(a.router)
}

func (a *App) Run() {
	a.router.Run()
}
