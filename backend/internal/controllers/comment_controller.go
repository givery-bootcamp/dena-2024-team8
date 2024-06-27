package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentCreate(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.PostForm("postId"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	body := ctx.PostForm("body")
	if body == "" {
		handleError(ctx, 400, errors.New("body is required"))
		return
	}

	userId, err := GetUserId(ctx)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewCommentUsecase(repository)
	result, err := usecase.Create(postId, body, userId)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	ctx.JSON(200, result)
}

func CommentUpdate(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	body := ctx.PostForm("body")
	if body == "" {
		handleError(ctx, 400, errors.New("body is required"))
		return
	}

	userId, err := GetUserId(ctx)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewCommentUsecase(repository)
	result, err := usecase.Update(commentId, body, userId)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	ctx.JSON(200, result)
}

func CommentDelete(ctx *gin.Context) {
	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	userId, err := GetUserId(ctx)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewCommentUsecase(repository)
	err = usecase.Delete(commentId, userId)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	ctx.JSON(204, nil)
}
