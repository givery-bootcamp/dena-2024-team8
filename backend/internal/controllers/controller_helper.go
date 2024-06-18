package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloErrorResponse struct {
	Message string "json:`message`"
}

func DB(ctx *gin.Context) *gorm.DB {
	return ctx.MustGet("db").(*gorm.DB)
}

func handleError(ctx *gin.Context, status int, err error) {
	res := HelloErrorResponse{
		Message: err.Error(),
	}
	ctx.JSON(status, &res)
}
