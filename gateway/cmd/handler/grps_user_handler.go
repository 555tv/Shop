package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	user, err := h.UserClient.GetUser(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(
			w,
			"User not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(
			w,
			"Failed to encode response",
			http.StatusInternalServerError,
		)
		return
	}
}
