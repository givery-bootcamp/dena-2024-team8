package main

import (
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/external"
	"myapp/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	external.SetupDB()
	// Initialize Elasticsearch
	external.InitElasticSearch()

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction())
	app.Use(middleware.Cors())
	// app.Use(middleware.Auth())
	middleware.SetupRoutes(app)
	err := app.Run(fmt.Sprintf("%s:%d", config.HostName, config.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
