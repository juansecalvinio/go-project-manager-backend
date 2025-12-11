package main

import (
	"go-project-manager-backend/internal/config"
	"go-project-manager-backend/internal/domain/services"
	"go-project-manager-backend/internal/infrastructure/database"
	"go-project-manager-backend/internal/infrastructure/repositories"
	"go-project-manager-backend/internal/interfaces/http/handlers"
	"go-project-manager-backend/internal/interfaces/http/middleware"
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	if err := config.LoadEnv(".env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Get MongoDB configuration from environment
	mongoURI := config.GetEnv("MONGODB_URI", "mongodb://localhost:27017")
	mongoDB := config.GetEnv("MONGODB_DATABASE", "project_manager")

	// Connect to MongoDB
	db, err := database.NewMongoDB(mongoURI, mongoDB)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Close()

	// Initialize MongoDB repositories
	userRepository := repositories.NewMongoUserRepository(db.Database)
	taskRepository := repositories.NewMongoTaskRepository(db.Database)

	// Initialize services
	userService := services.NewUserService(userRepository)
	taskService := services.NewTaskService(taskRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)

	// Protected routes
	mux.HandleFunc("GET /users/profile", middleware.AuthMiddleware(userHandler.GetProfile))

	mux.HandleFunc("POST /tasks", middleware.AuthMiddleware(taskHandler.CreateTask))
	mux.HandleFunc("GET /tasks", middleware.AuthMiddleware(taskHandler.GetTask))
	mux.HandleFunc("PUT /tasks", middleware.AuthMiddleware(taskHandler.UpdateTask))
	mux.HandleFunc("DELETE /tasks", middleware.AuthMiddleware(taskHandler.DeleteTask))
	mux.HandleFunc("POST /tasks/assign", middleware.AuthMiddleware(taskHandler.AssignToSprint))
	mux.HandleFunc("POST /tasks/backlog", middleware.AuthMiddleware(taskHandler.MoveToBacklog))
	mux.HandleFunc("GET /tasks/list", middleware.AuthMiddleware(taskHandler.ListTasks))
	mux.HandleFunc("GET /tasks/backlog", middleware.AuthMiddleware(taskHandler.ListBacklog))

	port := config.GetEnv("PORT", "8080")

	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
