package services

import (
	"time"

	"gate/internal/models"
)

func ConvertUserRequestToUser(req models.UserRequest) models.User {
	return models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: req.BirthDate,
		CreatedAt: time.Now(),
	}
}

func ConvertUserToResponse(user models.User) models.UserResponse {
	return models.UserResponse{
		ID:        user.ID.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
	}
}

func ConvertUsersToResponse(users []models.User) []models.UserResponse {
	response := make([]models.UserResponse, 0, len(users))

	for _, user := range users {
		response = append(
			response,
			ConvertUserToResponse(user),
		)
	}

	return response
}
