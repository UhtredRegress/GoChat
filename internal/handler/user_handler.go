package handler

import "github.com/gin-gonic/gin"

type Registrar interface {
	Register(r *gin.Engine)
}
