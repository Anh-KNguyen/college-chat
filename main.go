package main

import (
	"fmt"
	"log"
	"net/http"
)

// handle all requests to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// match the URL path hit with a defined function
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// kick off API
func main() {
	handleRequests()
}

// Article struct for title, description, and content
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// global Articles array to populate in main function to simulate a database
var Articles []Article
