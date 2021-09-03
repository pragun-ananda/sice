// Serve API requests - GET, POST, etc.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Entry struct {
	Id        int     `json:"id"`
	Title     string  `json:"name"`
	Desc      string  `json:"desc"`
	Rating    int     `json:"rating"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// Will be replaced with a database
var Entries []Entry
var Id int = 0

func populateEntries() {
	Entries = []Entry{
		Entry{Id: 1, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
		Entry{Id: 2, Title: "Test", Desc: "Test Desc", Rating: 0, Latitude: 0.0, Longitude: 0.0},
	}
}

func returnAllEntries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Entries)
}

func returnSingleEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, entry := range Entries {
		if entry.Id == key {
			json.NewEncoder(w).Encode(entry)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllEntries)
	http.HandleFunc("/entries/{}", returnSingleEntry)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	populateEntries()
	handleRequests()
}

/*
Resources:
	- Encryption: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
*/
