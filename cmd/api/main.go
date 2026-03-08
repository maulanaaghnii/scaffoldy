package main

import (
	"log"
	"net/http"
	"scaffoldy/internal/auth"
	"scaffoldy/internal/category"
	"scaffoldy/internal/item"
	"scaffoldy/internal/itemCategory"
	"scaffoldy/internal/task"
	"scaffoldy/pkg/middleware"
	"time"

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

	// API Routes Group
	api := r.Group("/api")
	{
		// 1. PUBLIC ROUTES (Tanpa Auth)
		auth.Register(api, db) // /api/login

		// 2. PROTECTED ROUTES (Pakai Auth)
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// category.Register(protected, db)
			item.Register(protected, db)
			itemCategory.Register(protected, db)
			category.Register(protected, db)
			task.Register(protected, db)
			// [SCAFFOLDY_INSERT_MARKER]
		}
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
