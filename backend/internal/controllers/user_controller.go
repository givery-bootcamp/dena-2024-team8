package controllers

import (
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ctxからuser_idを取得し、user_idに紐づくユーザー情報を取得して返す
func UserDetail(ctx *gin.Context) {
	//ctxからuserIdを取得
	userId, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}
	intUserId, ok := userId.(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is invalid"})
		return
	}

	//userIdに紐づくユーザー情報を取得
	repository := repositories.NewUserRepository(DB(ctx))
	usecase := usecases.NewUserUsecase(repository)
	result, err := usecase.Get(intUserId)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, result)   
	}
}