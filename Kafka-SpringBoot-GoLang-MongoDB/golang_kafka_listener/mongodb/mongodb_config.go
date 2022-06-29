package mongodb

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const dbConnectionUri = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func GetClientInstance() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConnectionUri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	coll := client.Database("kafka").Collection("message")

	return coll
}

func InsertMessage(collection *mongo.Collection, message string) {
	messageSlice := strings.Split(message, ":")
	doc := bson.D{{messageSlice[0], messageSlice[1]}}

	result, err := collection.InsertOne(context.TODO(), doc)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}
