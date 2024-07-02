package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection
var client *mongo.Client
var cred options.Credential

func Connection() *mongo.Collection {
	var err error

	cred.Username = os.Getenv("MONGO_USERNAME")
	cred.Password = os.Getenv("MONGO_PASSWORD")

	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(cred)

	// Connect to MongoDB
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("MONGO: ", err)
		return nil
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("MONGO: ", err)
		return nil
	}

	// Set the database and collection variables
	db = client.Database("teste_backend").Collection("products")

	return db
}

func Disconnect() {
	client.Disconnect(context.TODO())
}
