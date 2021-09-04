// Serve API requests - GET, POST, etc.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

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

var Entries []Entry
var idCount int64 = 0
var client mongo.Client

func populateEntries() {
	Entries = []Entry{
		Entry{Id: 1, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
		Entry{Id: 2, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
	}
}

func getAllEntries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Entries)
}

func getSingleEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	numId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range Entries {
		if entry.Id == numId {
			json.NewEncoder(w).Encode(entry)
		}
	}
}

func createNewEntry(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var entry Entry
	json.Unmarshal(reqBody, &entry)
	entry.Id = idCount
	Entries = append(Entries, entry)
	json.NewEncoder(w).Encode(entry)
	idCount++
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	numId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	for index, entry := range Entries {
		if entry.Id == numId {
			Entries = append(Entries[:index], Entries[index+1:]...)
		}
	}
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	deleteEntry(w, r)
	createNewEntry(w, r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func connectToDB() mongo.NewClient {
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

	collection := client.Database("sice").Collection("entries")
	return collection
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	go myRouter.HandleFunc("/", homePage)
	go myRouter.HandleFunc("/all", getAllEntries).Methods("GET")
	go myRouter.HandleFunc("/entries/{id}", getSingleEntry).Methods("GET")
	go myRouter.HandleFunc("/entries", createNewEntry).Methods("POST")
	go myRouter.HandleFunc("/entries/{id}", deleteEntry).Methods("DELETE")
	go myRouter.HandleFunc("/entries/{id}", updateEntry).Methods("UPDATE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	populateEntries()
	connectToDB()
	handleRequests()
}

/*
Resources:
	- Encryption: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
*/
