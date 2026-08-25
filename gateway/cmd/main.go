package main

import (
	"fmt"
	"log"
	"net/http"

	"gate/cmd/handler"
	"gate/internal/clients"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Подключаемся к User Service
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Создаём User Client
	userClient := clients.NewUserClient(conn)

	// Пока просто проверяем, что клиент создан
	_ = userClient

	// Регистрируем HTTP-маршруты
	handler.RegisterRoutes(userClient)

	// Подключаем папку со статикой
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Gateway started: http://localhost:8080")

	// Запускаем HTTP-сервер
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
