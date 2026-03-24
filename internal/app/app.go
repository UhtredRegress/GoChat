package app

import "github.com/gin-gonic/gin"

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	router := gin.Default()
	return &App{router}
}

func (a *App) Run() {

}
