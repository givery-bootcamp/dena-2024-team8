package controllers

import (
	"github.com/gin-gonic/gin"
)

func SignOut(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "signout success",
	})
}
