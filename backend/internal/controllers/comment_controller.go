package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentList(ctx *gin.Context) {
	postId, err := strconv.Atoi(ctx.Param("postId"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	repository := repositories.NewCommentRepository(DB(ctx))
	usecase := usecases.NewCommentUsecase(repository)
	result, err := usecase.List(postId, limit, offset)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}
	ctx.JSON(200, result)
}

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
