// Serve API requests - GET, POST, etc.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Entries []Entry
var idCount int64 = 0

func populateEntries() {
	Entries = []Entry{
		Entry{Id: 1, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
		Entry{Id: 2, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
	}
}

func getAllRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Entries)
}

func getRequest(w http.ResponseWriter, r *http.Request) {
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

func postRequest(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var entry Entry
	json.Unmarshal(reqBody, &entry)
	entry.Id = idCount
	Entries = append(Entries, entry)
	json.NewEncoder(w).Encode(entry)
	idCount++
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
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

func putRequest(w http.ResponseWriter, r *http.Request) {
	deleteRequest(w, r)
	postRequest(w, r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	go myRouter.HandleFunc("/", homePage)
	go myRouter.HandleFunc("/all", getAllRequest).Methods("GET")
	go myRouter.HandleFunc("/entries/{id}", getRequest).Methods("GET")
	go myRouter.HandleFunc("/entries", postRequest).Methods("POST")
	go myRouter.HandleFunc("/entries/{id}", deleteRequest).Methods("DELETE")
	go myRouter.HandleFunc("/entries/{id}", putRequest).Methods("UPDATE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	populateEntries()
	client = connectToDB() // source: https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver
	collection := client.Database("sice").Collection("entries")
	handleRequests()
}

/*
Resources:
	- Encryption: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

TODO:
	- abstract database CRUD operations (into a separate file) and keep API handling in this file
	- add security (API Key)

*/
