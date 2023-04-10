package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connect() *mongo.Client {
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_HOST"))
	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDatabase() *mongo.Database {
	return client.Database(os.Getenv("DB_NAME"))
}
