package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handle all requests to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// match the URL path hit with a defined function
func handleRequests() {
	// creates new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles) // add articles route and map to returnAllArticles function
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// kick off API
func main() {
	fmt.Println("REST API v2.0 - Mux Routers")
	// populate Articles with dummy data
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

// Article struct for title, description, and content
type Article struct {
	Id      string `json:"Id"`
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

// return single article based on {id} value from URL
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// loop through articles and return matched article as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}
