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
	/* Returns a JSON response encoding all records */
	fmt.Println("Endpoint Hit: getAllEntries")
	json.NewEncoder(w).Encode(Entries)
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	/* Returns JSON response for a specified item id*/
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
	/* Handles request to POST a new entry */
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
	/* Handles request to DELETE an entry with a specified id */
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
	/* Handles PUT request to update a record (deletes old record and posts new one) */
	deleteRequest(w, r)
	postRequest(w, r)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	/* Mock function to test API root endpoint */
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	/* Helper function to handle API routing */
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
	client = connectToDB()
	collection := client.Database("sice").Collection("entries")
	handleRequests()
}

/*
TODO:
	- add API security
*/
