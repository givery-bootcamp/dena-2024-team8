package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"

	"github.com/gin-gonic/gin"
)

func SignIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if username == "" || password == "" {
		ctx.JSON(400, gin.H{
			"error": "username and password are required",
		})
		return
	}

	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)
	result, jwtToken, err := usecase.VerifyUserAndGenerateJWT(username, password)
	ctx.SetCookie("jwt", *jwtToken, 3600, "/", "", false, true)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("signin failed"))
	}
}

func SignOut(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "signout success",
	})
}
