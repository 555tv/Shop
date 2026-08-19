package handler

import "net/http"

// RegisterRoutes регистрирует все HTTP-маршруты.
func RegisterRoutes() {

	h := &Handler{}

	// CRUD
	http.HandleFunc("GET /users", h.GetUsers)
	http.HandleFunc("POST /users", h.CreateUser)
	http.HandleFunc("PUT /users/{id}", h.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", h.DeleteUser)

	// Поиск пользователей
	http.HandleFunc("GET /users/search", h.FindUsers)
}
