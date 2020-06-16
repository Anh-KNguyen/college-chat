package main

import (
	"encoding/json"
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
	http.HandleFunc("/articles", returnAllArticles) // add articles route and map to returnAllArticles function
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// kick off API
func main() {
	// populate Articles with dummy data
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
	}
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

// create REST endpoint when hit with a HTTP GET request, will return all articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // encodes articles array into JSON string and write as part of response
}
