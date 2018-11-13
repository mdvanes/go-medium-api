// https://tour.golang.org/basics/3
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
// https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
// https://mholt.github.io/json-to-go/
// https://medium.com/statuscode/building-a-basic-web-service-to-display-your-medium-blog-posts-on-your-website-using-aws-api-48597b1771c5

// Install dependencies:
// go get github.com/gorilla/mux

// Building
// go build && ./api

// Path of this file
// ~/go/src/github.com/mdvanes/medium-api/api/main.go

package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
)

// our main function
func main() {
	fmt.Println("Server running, try GET on http://localhost:8000/posts")

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/", fs)
	router.Handle("/main.js", fs)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	url := "https://medium.com/codestar-blog/latest?format=json"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Strip security prefix
	if resp.StatusCode == http.StatusOK {
		var posts, users = ParseResponse(resp)
		LogResponse(posts)
		simplePosts := ConvertPosts(posts, users)
		json.NewEncoder(w).Encode(simplePosts)
	}
}

func ParseResponse(resp *http.Response) ([]Post, map[string]User) {
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	saneBodyString := bodyString[16:len(bodyString)]
	//log.Println(bodyString[16:len(bodyString)])

	// Fill the record with the data from the JSON
	var record LatestResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(strings.NewReader(saneBodyString)).Decode(&record); err != nil {
		log.Println(err)
	}

	return record.Payload.Posts, record.Payload.References.User
}

func LogResponse(posts []Post) {
	log.Println("Nr of posts =", len(posts))
	//fmt.Println("status =", record.Success)
	//fmt.Println("First Title =", posts[0].Title, posts[0].LatestPublishedAt, posts[0].PreviewContent.BodyModel.Paragraphs[1].Text)
	//log.Println("Users =", record.Payload.References)
	//for _, user := range users {
	//	log.Println("user =", user)
	//}
}
