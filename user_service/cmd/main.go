package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := pgx.ParseConfig(
		"postgres://user_service@127.0.0.1:5432/users_db?sslmode=disable",
	)

	if err != nil {
		log.Fatal("Config error:", err)
	}

	config.Password = "user_service_password"

	conn, err := pgx.ConnectConfig(
		context.Background(),
		config,
	)

	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	defer conn.Close(context.Background())

	log.Println("Connected to PostgreSQL")
}
