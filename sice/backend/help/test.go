// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Hi there, I love %s!", path)
}

func main() {
	go http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
