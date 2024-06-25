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
	"golang.org/x/crypto/bcrypt"
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

	testCases := []struct{ username, password string }{
		{"taro", "password"},
		{"hanako", "PASSWORD"},
	}

	for _, tc := range testCases {
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

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, tc.username, ru.Name)
		err := bcrypt.CompareHashAndPassword([]byte(ru.Password), []byte(tc.password))
		assert.NoError(t, err)
	}
}
