package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "myapp/docs"
	"myapp/internal/controllers"
)

// @title Example API
// @version 1.0
// @description This is a sample server for demonstrating Swagger with Gin.
// @host localhost:8080
// @BasePath /

func SetupRoutes(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "It works")
	})
	app.GET("/hello", controllers.HelloWorld)
	app.GET("/posts", controllers.PostList)
	app.GET("/posts/:postId", controllers.PostDetail)
	app.POST("/signin", controllers.SignIn)
	// Swaggerのエンドポイントを設定
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := app.Group("/auth", Auth())
	auth.GET("/user", controllers.UserDetail)
	// Userのエンドポイントを設定
}
