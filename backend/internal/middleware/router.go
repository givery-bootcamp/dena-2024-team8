package middleware

import (
	"myapp/internal/controllers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "path/to/your/project/docs"
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
	// Swaggerのエンドポイントを設定
    app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
