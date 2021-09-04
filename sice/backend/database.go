// Connect to MongoDB Driver - Manage CRUD Operations

package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Entry struct {
	Id        int64   `json:"id"`
	Title     string  `json:"name"`
	Desc      string  `json:"desc"`
	Rating    int     `json:"rating"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

var client *mongo.Client
var collection *mongo.Collection

func connectToDB() *mongo.Client {
	// parse mongoauth.json and replace Mongo URI with it
	// source: https://www.mongodb.com/blog/post/quick-start-golang-mongodb-starting-and-setup
	client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI_HERE>"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func create(newEntry Entry) primitive.ObjectID {
	insertResult, err := collection.InsertOne(context.TODO(), newEntry)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID // what type is this??
}

func read(findEntry Entry, filter bson.D) Entry {
	err := collection.FindOne(context.TODO(), filter).Decode(&findEntry)
	if err != nil {
		log.Fatal(err)
	}
	return findEntry
}

func update(filter bson.D, updateTerms bson.D) (int64, int64) {
	updateResult, err := collection.UpdateOne(context.TODO(), filter, updateTerms)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult.MatchedCount, updateResult.ModifiedCount
}

func delete(filter bson.D) int64 {
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}

// func main() {
// 	fmt.Println("Test")
// }

/*
https://github.com/tfogo/mongodb-go-tutorial/blob/master/main.go
*/
