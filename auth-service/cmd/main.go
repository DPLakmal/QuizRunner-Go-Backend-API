package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/auth-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/auth-service/domain"
	"github.com/pubudulakmal/quiz-backend/auth-service/repository"
	"github.com/pubudulakmal/quiz-backend/auth-service/usecase"
	"github.com/pubudulakmal/quiz-backend/pkg/db"
)

func main() {
	database := db.InitDB()
	defer database.Close()

	// Auto migrate
	database.AutoMigrate(&domain.User{})

	r := gin.Default()

	authRepo := repository.NewPostgresAuthRepository(database)
	authUC := usecase.NewAuthUseCase(authRepo)

	// Load config from mounted volume or use environment variable
	config := domain.LoadConfig("/configs/auth.json")
	if config == nil {
		config = auth.Config
	}
	http.NewAuthHandler(r, authUC, config.JWTSecret)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Auth Service starting on port %s", port)
	log.Printf("Using JWT_SECRET configured: %s", config.JWTSecret)
	r.Run(":" + port)
}
