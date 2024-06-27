package middleware

import (
	_ "myapp/docs"
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"      // swagger embed files
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
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
	app.GET("/posts/search", controllers.PostSearch)
	app.GET("/posts/:postId", controllers.PostDetail)
	app.POST("/signin", controllers.SignIn)
	app.POST("/signout", controllers.SignOut)
	// Swaggerのエンドポイントを設定
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := app.Group("/auth", Auth())
	auth.GET("/user", controllers.UserDetail)
	// Userのエンドポイントを設定
}
