package controllers

import (
	"errors"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//swaggo用の型定義
type Post struct {
	Id        int       `json:"id"`
	UserId    int       `json:user_id`
	Title     string    `json:title`
	Body      string    `json:body`
	CreatedAt time.Time `json:created_at`
	UpdatedAt time.Time `json:update_at`
	DeletedAt time.Time `json:deleted_at`
}

// PostList godoc
// @Summary get post list
// @Description get post list
// @ID get-post-list
// @Tags post
// @Accept  json
// @Produce  json
// @Success 200 {array} Post
// @Router /posts [get]
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

// PostDetail godoc
// @Summary get post by id
// @Description get post by id
// @ID get-post-by-id
// @Tags post
// @Accept  json
// @Produce  json
// @Param postId path int true "Post ID"
// @Success 200 {object} Post[]
// @Router /posts/{postId} [get]
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
