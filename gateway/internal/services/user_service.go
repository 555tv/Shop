package services

import (
	"errors"

	"gate/internal/models"
	"gate/internal/storage"
	"gate/internal/validator"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Получить всех пользователей.
func GetUsers() ([]models.UserResponse, error) {

	users, err := storage.GetUsers()
	if err != nil {
		return nil, err
	}

	return ConvertUsersToResponse(users), nil
}

// Получить пользователя по ID.
func GetUserByID(id string) (*models.UserResponse, error) {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user, err := storage.GetUserByID(objectID)
	if err != nil {
		return nil, err
	}

	response := ConvertUserToResponse(*user)

	return &response, nil
}

// Создать пользователя.
func CreateUser(request models.UserRequest) error {

	// Проверяем данные
	if err := validator.Validate(request); err != nil {
		return err
	}

	// Проверяем существование пользователя
	exists, err := storage.UserExists(
		request.FirstName,
		request.LastName,
	)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	}

	// Конвертируем запрос в модель
	user := ConvertUserRequestToUser(request)

	// Сохраняем
	return storage.InsertUser(user)
}

// Обновить пользователя.
func UpdateUser(
	id string,
	request models.UpdateUserRequest,
) error {

	if err := validator.Validate(request); err != nil {
		return err
	}

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return storage.UpdateUser(
		objectID,
		request,
	)
}

// Удалить пользователя.
func DeleteUser(id string) error {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	return storage.DeleteUser(objectID)
}

// Поиск пользователей.
func FindUsers(
	firstName string,
	lastName string,
) ([]models.UserResponse, error) {

	users, err := storage.FindUsersByName(
		firstName,
		lastName,
	)

	if err != nil {
		return nil, err
	}

	return ConvertUsersToResponse(users), nil
}
