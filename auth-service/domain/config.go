package domain

import (
	"encoding/json"
	"os"
)

// AuthConfig holds configuration for auth service
type AuthConfig struct {
	JWTSecret string `json:"jwtSecret"`
	Port      string `json:"port"`
}

// LoadConfig loads configuration from a JSON file
// If the file doesn't exist or is invalid, return nil
// The JWT_SECRET environment variable should be used as fallback
func LoadConfig(path string) *AuthConfig {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var config AuthConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil
	}
	return &config
}
