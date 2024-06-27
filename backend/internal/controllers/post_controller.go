package controllers

import (
	"errors"
	"myapp/internal/entities"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
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

// PostCreate godoc
// @Summary create post
// @Description create post
// @ID create-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param title query string true "タイトル (最大100文字)"
// @Param body query string true "本文"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse "タイトルと本文は必須です。"
// @Failure 400 {object} ErrorResponse "タイトルは100文字以下である必要があります。"
// @Failure 401 {object} ErrorResponse "認証が必要です。"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /posts [post]
func PostCreate(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	title := ctx.PostForm("title")
	body := ctx.PostForm("body")

	// ユーザーIDが取得できない場合は認証エラー
	if userId == 0 {
		handleError(ctx, 401, errors.New("authentication required"))
		return
	}

	// タイトルと本文は必須
	if title == "" || body == "" {
		handleError(ctx, 400, errors.New("title and body are required"))
		return
	}

	// タイトルは100文字以下
	if len(title) > 100 {
		handleError(ctx, 400, errors.New("title is too long"))
		return
	}

	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	result, err := usecase.Create(title, body, userId)
	if err != nil {
		handleError(ctx, 500, err)
	} else {
		ctx.JSON(200, result)
	}
}

// backend
// endpoint: PUT /posts/:postId
// parameters: JSON
// title: string(必須。100文字以下)
// body: string(必須)
// 認証: 必要
// response:
// 200 更新されたPostエンティティ
// 400 対象の投稿の作成者が自分ではない場合

// PostUpdate godoc
// @Summary update post
// @Description update post
// @ID update-post
// @Tags post
// @Accept  json
// @Produce  json
// @Param postId path int true "Post ID"
// @Param title query string true "タイトル (最大100文字)"
// @Param body query string true "本文"
// @Success 200 {object} Post
// @Failure 400 {object} ErrorResponse "タイトルと本文は必須です。"
// @Failure 400 {object} ErrorResponse "タイトルは100文字以下である必要があります。"
// @Failure 401 {object} ErrorResponse "認証が必要です。"
// @Failure 400 {object} ErrorResponse "対象の投稿の作成者が自分ではない場合"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /posts/{postId} [put]
func PostUpdate(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	sid := ctx.Param("postId")
	id, err := strconv.Atoi(sid)
	if err != nil {
		handleError(ctx, 500, err)
	}

	title := ctx.PostForm("title")
	body := ctx.PostForm("body")

	// ユーザーIDが取得できない場合は認証エラー
	if userId == 0 {
		handleError(ctx, 401, errors.New("authentication required"))
		return
	}

	// タイトルと本文は必須
	if title == "" || body == "" {
		handleError(ctx, 400, errors.New("title and body are required"))
		return
	}

	// タイトルは100文字以下
	if len(title) > 100 {
		handleError(ctx, 400, errors.New("title is too long"))
		return
	}

	repository := repositories.NewPostRepository(DB(ctx))
	usecase := usecases.NewPostUsecase(repository)
	result, err := usecase.Update(title, body, userId, id)
	if err != nil {
		handleError(ctx, 500, err)
	}
	if result == nil {
		handleError(ctx, 400, errors.New("you are not the author of this post"))
	} else {
		ctx.JSON(200, result)
	}
}
