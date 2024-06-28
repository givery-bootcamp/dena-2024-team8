package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HelloErrorResponse struct {
	Message string `json:"message"`
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

func GetUserId(ctx *gin.Context) (int, error) {
	userId, exists := ctx.Get("userId")
	if !exists {
		return 0, fmt.Errorf("user_id is required")
	}
	intUserId, ok := userId.(int)
	if !ok {
		return 0, fmt.Errorf("user_id is invalid")
	}
	return intUserId, nil
}
