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

	// Initialize Item dependencies
	itemRepo := item.NewRepository(db)
	itemService := item.NewService(itemRepo)
	itemHandler := item.NewHandler(itemService)

	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := category.NewHandler(categoryService)

	// Initialize Category dependencies
	// categoryRepo := category.NewRepository(db)
	// categoryService := category.NewService(categoryRepo)
	// categoryHandler := category.NewHandler(categoryService)

	// Initialize User dependencies
	// userRepo := user.NewRepository(db)
	// userService := user.NewService(userRepo)
	// userHandler := user.NewHandler(userService)

	// Setup Gin router
	r := gin.Default()

	// Set Cors
	if cfg.Environment == "development" {
		// Development: Allow all origins
		r.Use(cors.Default())
		log.Println("CORS: Allowing all origins (Development mode)")
	} else {
		// Production: Restrict origins
		r.Use(cors.New(cors.Config{
			AllowOrigins:     cfg.AllowedOrigins,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
		log.Printf("CORS: Allowing origins: %v (Production mode)", cfg.AllowedOrigins)
	}
	// API Routes Group
	api := r.Group("/api")
	{
		// Item routes
		items := api.Group("/items")
		{
			items.GET("", itemHandler.GetAllItem)
			items.POST("", itemHandler.CreateItem)
			items.GET("/:id", itemHandler.GetItemByID)
			items.GET("/code/:code", itemHandler.GetItemByCode)
			items.PUT("/:id", itemHandler.UpdateItem)
			items.DELETE("/:id", itemHandler.SoftDeleteItem)
		}

		// Category routes
		categories := api.Group("/category")
		{
			categories.GET("", categoryHandler.GetAllCategory)
			categories.POST("", categoryHandler.CreateCategory)
			categories.GET("/:id", categoryHandler.GetCategoryByID)
			categories.GET("/code/:code", categoryHandler.GetCategoryByCode)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.SoftDeleteCategory)
		}

		// // Category routes
		// categories := api.Group("/category")
		// {
		// 	categories.GET("", categoryHandler.GetAllCategory)
		// 	categories.POST("", categoryHandler.CreateCategory)
		// 	categories.GET("/:id", categoryHandler.GetCategoryByID)
		// 	categories.GET("/code/:code", categoryHandler.GetCategoryByCode)
		// 	categories.PUT("/:id", categoryHandler.UpdateCategory)
		// 	categories.DELETE("/:id", categoryHandler.DeleteCategory)
		// }

		// // User routes
		// api.POST("/auth/login", userHandler.Login)
		// users := api.Group("/users")
		// {
		// 	users.GET("", userHandler.GetUsers)
		// 	users.POST("", userHandler.CreateUser)
		// 	users.GET("/:id", userHandler.GetUser)
		// 	users.GET("/username/:username", userHandler.GetUserByUsername)
		// 	users.PUT("/:id", userHandler.UpdateUser)
		// 	users.POST("/:id/change-password", userHandler.ChangePassword)
		// 	users.DELETE("/:id", userHandler.DeleteUser)
		// }
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
