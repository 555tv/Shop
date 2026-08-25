package handler

import (
	"net/http"

	"gate/internal/clients"
)

func RegisterRoutes(userClient *clients.UserClient) {

	h := &Handler{
		UserClient: userClient,
	}
	http.HandleFunc("GET /users/{id}", h.GetUser)
	// CRUD
	http.HandleFunc("GET /users", h.GetUsers)
	http.HandleFunc("POST /users", h.CreateUser)
	http.HandleFunc("PUT /users/{id}", h.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", h.DeleteUser)

	// Поиск пользователей
	http.HandleFunc("GET /users/search", h.FindUsers)

}
