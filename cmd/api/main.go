package main

import (
	"log"
	"net/http"
	"time"

	"scaffoldy/internal/category"
	"scaffoldy/internal/item"

	"scaffoldy/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.Load()

	// Connect to database
	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to MariaDB")

	// Setup Gin router
	r := gin.Default()

	// Set Cors
	if cfg.Environment == "development" {
		r.Use(cors.Default())
		log.Println("CORS: Allowing all origins (Development mode)")
	} else {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	}

	// Dynamic Registration
	api := r.Group("/api")
	{
		item.Register(api, db)
		category.Register(api, db)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "up",
			"environment": cfg.Environment,
		})
	})

	// Start server
	serverAddr := ":" + cfg.ServerPort
	log.Printf("Server running on %s (Environment: %s)", serverAddr, cfg.Environment)
	if err := r.Run(serverAddr); err != nil {
		log.Fatal(err)
	}
}
