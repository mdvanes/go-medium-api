// https://tour.golang.org/basics/3
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo

// Install dependencies:
// go get github.com/gorilla/mux

// Building
// go build && ./api

// Path of this file
// ~/go/src/github.com/mdvanes/medium-api/api/main.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

// our main function
func main() {
	articles = append(articles, Article{ID: "1", Title: "My Title"})

	fmt.Println("Server running, try GET on http://localhost:8000/articles")

	router := mux.NewRouter()
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

type Article struct {
	ID        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
}

var articles []Article
