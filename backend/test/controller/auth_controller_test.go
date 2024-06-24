package repositories

import (
	"myapp/internal/controllers"
	"myapp/internal/external"
	"myapp/internal/interfaces"
	"myapp/internal/middleware"
	"myapp/internal/repositories"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func Init() {
	external.SetupDB()
}

func setupUser() (interfaces.UserRepository, func()) {
	db := external.DB.Begin()
	repo := repositories.NewUserRepository(db)
	teardown := func() {
		db.Rollback()
	}

	// Insert test data
	db.Create(&repositories.User{
		Name:     "user1",
		Password: "password1",
	})
	db.Create(&repositories.User{
		Name:     "user2",
		Password: "password2",
	})
	return repo, teardown
}

func TestSignIn(t *testing.T) {
	// Setup webserver
	// app := gin.Default()
	// app.Use(middleware.Transaction())
	// app.Use(middleware.Cors())
	// middleware.SetupRoutes(app)
	// app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))

	// Create a new Gin context for testing
	ctx, testApp := gin.CreateTestContext(httptest.NewRecorder())
	testApp.Use(middleware.Transaction())
	testApp.Use(middleware.Cors())
	middleware.SetupRoutes(testApp)

	// Set the request form values
	ctx.Request = httptest.NewRequest(http.MethodPost, "/signin", nil)
	ctx.Request.PostForm = url.Values{
		"username": {"user1"},
		"password": {"password1"},
	}

	controllers.SignIn(ctx)
}
