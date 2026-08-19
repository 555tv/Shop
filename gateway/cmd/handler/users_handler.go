package handler

import (
	"encoding/json"
	"net/http"

	"gate/internal/models"
	"gate/internal/services"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var request models.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(request); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("User created"))
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	id := r.PathValue("id")

	var request models.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.UpdateUser(id, request); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User updated"))
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	if err := services.DeleteUser(id); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("User deleted"))
}

func (h *Handler) FindUsers(w http.ResponseWriter, r *http.Request) {

	firstName := r.URL.Query().Get("firstName")
	lastName := r.URL.Query().Get("lastName")

	users, err := services.FindUsers(
		firstName,
		lastName,
	)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
