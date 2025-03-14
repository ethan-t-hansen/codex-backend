package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ethan-t-hansen/codex-backend/config"
	"github.com/ethan-t-hansen/codex-backend/handlers"
	"github.com/ethan-t-hansen/codex-backend/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	newsService := services.NewNewsService(cfg.NewsAPIKey)

	// Initialize handlers
	newsHandler := handlers.NewNewsHandler(newsService)

	// Initialize router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Define API endpoints
	api := router.Group("/api")
	{
		news := api.Group("/news")
		{
			news.GET("/top-headlines", newsHandler.GetTopHeadlines)
			news.GET("/search", newsHandler.SearchNews)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	router.Run(":" + cfg.Port)
}