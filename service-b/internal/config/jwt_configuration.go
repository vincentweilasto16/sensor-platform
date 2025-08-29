package config

import (
	"log"
	"os"
	"strconv"
)

type JWTConfig struct {
	Secret    string
	ExpiresIn int64 // in seconds
}

func LoadJWTConfig() *JWTConfig {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET must be set in environment")
	}

	expiresStr := os.Getenv("JWT_EXPIRES_IN")
	if expiresStr == "" {
		log.Fatal("JWT_EXPIRES_IN must be set in environment")
	}

	expires, err := strconv.ParseInt(expiresStr, 10, 64)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRES_IN: %v", err)
	}

	return &JWTConfig{
		Secret:    secret,
		ExpiresIn: expires,
	}
}
