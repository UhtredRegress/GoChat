package main

import (
	"github.com/UhtredRegress/GoChat/internal/app"
)

func main() {
	application := app.NewApp()
	application.RegisterHandler()
	application.Run()
}
