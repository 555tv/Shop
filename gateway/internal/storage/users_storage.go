package storage

import (
	"context"
	"errors"

	"gate/internal/models"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func InsertUser(user models.User) error {

	_, err := collection.InsertOne(
		context.Background(),
		user,
	)

	return err
}

func GetUsers() ([]models.User, error) {

	cursor, err := collection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var users []models.User

	err = cursor.All(
		context.Background(),
		&users,
	)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// ===============================
// READ BY ID
// ===============================

func GetUserByID(id bson.ObjectID) (*models.User, error) {

	var user models.User

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// ===============================
// UPDATE
// ===============================

func UpdateUser(
	id bson.ObjectID,
	user models.UpdateUserRequest,
) error {

	update := bson.M{
		"$set": bson.M{
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"birthDate": user.BirthDate,
		},
	}

	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// ===============================
// DELETE
// ===============================

func DeleteUser(id bson.ObjectID) error {

	result, err := collection.DeleteOne(
		context.Background(),
		bson.M{"_id": id},
	)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// ===============================
// SEARCH
// ===============================

func FindUsersByName(
	firstName string,
	lastName string,
) ([]models.User, error) {

	filter := bson.M{}

	if firstName != "" {
		filter["firstName"] = firstName
	}

	if lastName != "" {
		filter["lastName"] = lastName
	}

	cursor, err := collection.Find(
		context.Background(),
		filter,
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var users []models.User

	err = cursor.All(
		context.Background(),
		&users,
	)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// ===============================
// EXISTS
// ===============================

func UserExists(
	firstName string,
	lastName string,
) (bool, error) {

	count, err := collection.CountDocuments(
		context.Background(),
		bson.M{
			"firstName": firstName,
			"lastName":  lastName,
		},
	)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
