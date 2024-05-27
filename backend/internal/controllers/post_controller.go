package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
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

func PostDetail(ctx *gin.Context) {
	sid := ctx.Param("postId")
	id, err := strconv.Atoi(sid)
	if err != nil {
		handleError(ctx, 500, err)
	}

	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	result, err := usecase.Get(id)
	if err != nil {
		handleError(ctx, 500, err)
	} else if result != nil {
		ctx.JSON(200, result)
	} else {
		handleError(ctx, 404, errors.New("not found"))
	}
}
