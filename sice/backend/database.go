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
	/*
		Uses secure URI to connect to MongoDB Atlas via MongoDB Go Driver and returns client connection

		Source: https://www.mongodb.com/blog/post/quick-start-golang-mongodb-starting-and-setup
		Source: https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
	*/
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
	/* Inserts a new entry and returns the inserted record id */
	insertResult, err := collection.InsertOne(context.TODO(), newEntry)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult.InsertedID
}

func read(findEntry Entry, filter bson.D) Entry {
	/* Finds an entry based on a BSON filter. Decodes the found entry and returns it or returns an error*/
	err := collection.FindOne(context.TODO(), filter).Decode(&findEntry)
	if err != nil {
		log.Fatal(err)
	}
	return findEntry
}

func update(filter bson.D, updateTerms bson.D) (int64, int64) {
	/* Finds an entry using a BSON filter and then updates it with BSON update terms */
	updateResult, err := collection.UpdateOne(context.TODO(), filter, updateTerms)
	if err != nil {
		log.Fatal(err)
	}
	return updateResult.MatchedCount, updateResult.ModifiedCount
}

func delete(filter bson.D) int64 {
	/* Deletes a record matching a BSON filter and returns the deleted record id*/
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount
}
