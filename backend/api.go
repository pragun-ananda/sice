// Serve API requests - GET, POST, etc.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func getAllEntries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Entries)
}

func getSingleEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, entry := range Entries {
		if entry.Id == key {
			json.NewEncoder(w).Encode(entry)
		}
	}
}

func createNewEntry(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var entry Entry
	json.Unmarshal(reqBody, &entry)
	Entries = append(Entries, entry)
	json.NewEncoder(w).Encode(entry)
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, entry := range Entries {
		if entry.Id == id {
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

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", getAllEntries).Methods("GET")
	myRouter.HandleFunc("/entries/{id}", getSingleEntry).Methods("GET")
	myRouter.HandleFunc("/entries", createNewEntry).Methods("POST")
	myRouter.HandleFunc("/entries/{id}", deleteEntry).Methods("DELETE")
	myRouter.HandleFunc("/entries/{id}", updateEntry).Methods("UPDATE")
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
