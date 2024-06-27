package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"myapp/internal/entities"
	"myapp/internal/external"
	"myapp/internal/repositories"
	"myapp/internal/usecases"
	"strconv"
	"strings"
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

func PostSearch(c *gin.Context) {
	query := c.Query("q")

	var searchQuery string
	if query == "" {
		// Query is empty, search for all documents
		searchQuery = `{
			"query": {
				"match_all": {}
			}
		}`
	} else {
		searchQuery = fmt.Sprintf(`{
			"query": {
				"multi_match": {
					"query": "%s",
					"fields": ["title", "body"]
				}
			}
		}`, query)
	}

	var buf strings.Builder
	buf.WriteString(searchQuery)

	res, err := external.ES.Search(
		external.ES.Search.WithContext(context.Background()),
		external.ES.Search.WithIndex("posts"),
		external.ES.Search.WithBody(strings.NewReader(buf.String())),
		external.ES.Search.WithTrackTotalHits(true),
		external.ES.Search.WithPretty(),
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Error parsing the response body: %s", err)})
			return
		} else {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Error: %s: %s", res.Status(), e["error"].(map[string]interface{})["reason"])})
			return
		}
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error parsing the response body: %s", err)})
		return
	}

	posts := make([]Post, 0)

	hits, ok := r["hits"].(map[string]interface{})
	if !ok || hits["hits"] == nil {
		// No hits found, handle empty result case
		c.JSON(200, posts)
		return
	}

	for _, hit := range hits["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		id, _ := source["id"].(float64) // Type assertion with fallback
		post := Post{
			Id:    int(id),
			Title: source["title"].(string),
			Body:  source["body"].(string),
		}
		posts = append(posts, post)
	}

	c.JSON(200, posts)
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
