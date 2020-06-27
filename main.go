package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET") // add articles route and map to returnAllArticles function
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST") //needs to be defined before the other /article endpoint (ordering)
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(myRouter)))
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
	fmt.Println("Endpoint hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	// loop through articles and return matched article as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")

	// get body of POST request, unmarshal into an Article struct and append to array
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")

	// parse path parameters
	vars := mux.Vars(r)
	id := vars["id"] // extract ID of article

	// search through articles, remove article if there is a match
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	vars := mux.Vars(r)
	id := vars["id"]

	for index, a := range Articles {
		if a.Id == id {
			Articles[index] = article
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))

}
