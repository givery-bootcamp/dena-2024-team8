package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// swaggo用の型定義
type Post struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// PostList godoc
// @Summary get post list
// @Description get post list
// @ID get-post-list
// @Tags post
// @Accept  json
// @Produce  json
// @Param limit query string false "件数 未実装です。"
// @Param offset query string false "開始位置 未実装です。"
// @Success 200 {array} Post
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /posts [get]
func PostList(ctx *gin.Context) {
	slimit := ctx.DefaultQuery("limit", "10")
	soffset := ctx.DefaultQuery("offset", "0")
	limit, err := strconv.Atoi(slimit)
	if err != nil {
		handleError(ctx, 500, err)
	}
	offset, err := strconv.Atoi(soffset)
	if err != nil {
		handleError(ctx, 500, err)
	}

	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	result, err := usecase.GetList(limit, offset)
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		if len(result) == 0 {
			result = []*entities.Post{}
		}
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
// @Param postId path int true "Post ID デフォルトで1から2までしかデータがありません。"
// @Success 200 {object} Post[]
// @Failure 400 {object} ErrorResponse "不正なpostID"
// @Failure 404 {object} ErrorResponse "ポストが見つからない"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
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


// PostDelete godoc
// @Summary delete post by id
// @Description delete post by id
// @ID delete-post-by-id
// @Tags post
// @Accept  json
// @Produce  json
// @Param postId path int true "Post ID デフォルトで1から2までしかデータがありません。"
// @Success 200
// @Failure 400 {object} ErrorResponse "不正なpostID"
// @Failure 404 {object} ErrorResponse "ポストが見つからない"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /posts/{postId} [delete]
func PostDelete(ctx *gin.Context) {
	sid := ctx.Param("postId")
	id, err := strconv.Atoi(sid)
	if err != nil {
		handleError(ctx, 500, err)
	}

	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	err = usecase.Delete(id)
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.Status(http.StatusOK)
	}
}