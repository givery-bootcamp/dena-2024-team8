package controllers

import (
	"github.com/gin-gonic/gin"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
)

func PostList(ctx *gin.Context) {
	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	result, err := usecase.GetList()
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, result)
	}
}
