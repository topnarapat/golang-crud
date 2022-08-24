package configs

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context
var Client *mongo.Client
var Collection *mongo.Collection

func MongoConnection() {
	// Create a new client and connect to the server
	context := context.Background()
	client, err := mongo.Connect(context, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect MongoDB Server Success")

	Ctx = context
	Client = client
}

func GetMongoCollection(collection string) *mongo.Collection {
	coll := Client.Database(os.Getenv("MONGO_DATABASE")).Collection(collection)
	return coll
}
