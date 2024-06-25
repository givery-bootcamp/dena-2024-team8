package repositories

import (
	"encoding/json"
	"myapp/internal/controllers"
	"myapp/internal/entities"
	"myapp/internal/external"
	"myapp/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	external.SetupDB()
}

func TestSignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Ginエンジンの初期化
	r := gin.New()
	r.Use(middleware.Transaction())
	r.Use(middleware.Cors())
	r.POST("/signin", controllers.SignIn)

	// テストリクエストの作成
	// user passwordをbodyに含める
	req, err := http.NewRequest(http.MethodPost, "/signin", nil)
	assert.NoError(t, err)

	type TestCase struct {
		name           string
		expectResponse int
		username       string
		password       string
	}

	testCases := []TestCase{
		{
			name:           "valid user1",
			expectResponse: http.StatusOK,
			username:       "taro",
			password:       "password",
		},
		{
			name:           "valid user2",
			expectResponse: http.StatusOK,
			username:       "hanako",
			password:       "PASSWORD",
		},
		{
			name:           "invalid user1",
			expectResponse: http.StatusBadRequest,
			username:       "",
			password:       "password",
		},
		{
			name:           "invalid user2",
			expectResponse: http.StatusBadRequest,
			username:       "taro",
			password:       "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// テストリクエストにbodyを追加
			req.PostForm = make(map[string][]string)
			req.PostForm.Add("username", tc.username)
			req.PostForm.Add("password", tc.password)

			// レスポンスのレコーダーを作成
			w := httptest.NewRecorder()

			// テストリクエストを実行
			r.ServeHTTP(w, req)

			ru := entities.User{}
			err = json.Unmarshal(w.Body.Bytes(), &ru)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.expectResponse, w.Code)
			if tc.expectResponse == http.StatusOK {
				assert.Equal(t, tc.username, ru.Name)
			}
		})
	}
}
