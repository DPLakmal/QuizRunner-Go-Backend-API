package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/auth-service/domain"
)

// defaultConfig is used when the config file cannot be loaded
var defaultConfig = &AuthConfig{
	JWTSecret: "supersecretkey", // Replace with a cryptographically secure random string in production
}

type AuthConfig struct {
	JWTSecret string
	Port      string
}

type AuthHandler struct {
	authUseCase domain.AuthUseCase
	config      *AuthConfig
}

func NewAuthHandler(r *gin.Engine, us domain.AuthUseCase, jwtSecret string) {
	// Load configuration from mounted volume
	config := loadConfig("/configs/auth.json")
	if config.JWTSecret == "" {
		config = defaultConfig
	}

	handler := &AuthHandler{
		authUseCase: us,
		config:      &config,
	}
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authUseCase.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUseCase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// loadConfig loads configuration from a JSON file
func loadConfig(path string) *AuthConfig {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		// Return default config if file doesn't exist
		return defaultConfig
	}
	var config AuthConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		// Return default config if JSON is invalid
		return defaultConfig
	}
	return &config
}
