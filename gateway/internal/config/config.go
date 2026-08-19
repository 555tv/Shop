package config

import (
	"os"
)

type Config struct {
	HTTPPort string

	UserServiceAddr    string
	ProductServiceAddr string
	OrderServiceAddr   string
	BillingServiceAddr string

	JWTSecret string
}

func Load() Config {
	return Config{
		HTTPPort:           getEnv("HTTP_PORT", "8080"),
		UserServiceAddr:    getEnv("USER_SERVICE_ADDR", "localhost:50051"),
		ProductServiceAddr: getEnv("PRODUCT_SERVICE_ADDR", "localhost:50052"),
		OrderServiceAddr:   getEnv("ORDER_SERVICE_ADDR", "localhost:50053"),
		BillingServiceAddr: getEnv("BILLING_SERVICE_ADDR", "localhost:50054"),
		JWTSecret:          getEnv("JWT_SECRET", "change-me"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
