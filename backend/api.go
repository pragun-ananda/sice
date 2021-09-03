// Serve API requests - GET, POST, etc.

package main

import (
	"fmt"
	"log"
	"net/http"
	// "encoding/json"
	// "github.com/gorilla/mux"
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
	Entries = []Entry {
		Entry{Id=1 }
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
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