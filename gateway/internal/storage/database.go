package storage

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

// ConnectMongo создает подключение к MongoDB.
func ConnectMongo() (*mongo.Client, error) {
	client, err := mongo.Connect(
		options.Client().ApplyURI(
			"mongodb://admin:password@localhost:27017",
		),
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}

// SetCollection сохраняет коллекцию,
// с которой будут работать CRUD-функции.
func SetCollection(c *mongo.Collection) {
	collection = c
}
