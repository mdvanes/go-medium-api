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
	//articles = append(articles, Article{ID: "1", Title: "My Title"})

	fmt.Println("Server running, try GET on http://localhost:8000/posts")

	router := mux.NewRouter()
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	url := "https://medium.com/codestar-blog/latest?format=json"

	// TODO split into smaller functions

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

		posts = record.Payload.Posts
		var users = record.Payload.References.User

		log.Println("Nr of posts =", len(posts))
		//fmt.Println("status =", record.Success)
		//fmt.Println("First Title =", posts[0].Title, posts[0].LatestPublishedAt, posts[0].PreviewContent.BodyModel.Paragraphs[1].Text)
		//log.Println("Users =", record.Payload.References)
		//for _, user := range users {
		//	log.Println("user =", user)
		//}

		// Write the struct to the response
		//json.NewEncoder(w).Encode(posts)

		// Convert Post to SimplePost
		var simplePosts []SimplePost

		// TODO use mapping
		// range posts returns index, item
		for _, item := range posts {
			var simplePost SimplePost
			simplePost.ID = item.ID
			simplePost.Title = item.Title
			simplePost.Author = users[item.CreatorId].Name
			simplePost.LatestPublishedAt = item.LatestPublishedAt
			simplePost.Paragraphs = item.PreviewContent.BodyModel.Paragraphs
			simplePosts = append(simplePosts, simplePost)
		}
		json.NewEncoder(w).Encode(simplePosts)
	}
}

type SimplePost struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	Author                 string `json:"author"`
	LatestPublishedAt      int    `json:"latestPublishedAt"`
	Paragraphs []struct {
		Text               string `json:"text"`
	} `json:"paragraphs"`
}

type Post struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	CreatorId              string `json:"creatorId"`
	LatestPublishedAt      int    `json:"latestPublishedAt"`
	PreviewContent struct {
		BodyModel struct {
			Paragraphs []struct {
				Text       string `json:"text"`
			} `json:"paragraphs"`
		} `json:"bodyModel"`
	} `json:"previewContent"`
}

type User struct {
	UserID            string `json:"userId"`              // userId (equal to Post.creatorId)
	Name              string `json:"name"`                // name (author name)
	Username          string `json:"username"`
	ImageID           string `json:"imageId"`
	BackgroundImageID string `json:"backgroundImageId"`
	Bio               string `json:"bio"`
	TwitterScreenName string `json:"twitterScreenName"`
}

// Auto generated with https://mholt.github.io/json-to-go/ from https://medium.com/codestar-blog/latest?format=json
type LatestResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Posts []Post `json:"posts,omitempty"`
		References struct {
			User map[string]User
		}
	} `json:"payload"`
}

var posts []Post
