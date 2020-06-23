package main

/*

Creating a Simple RESTful API in Go
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, re *http.Request) {
	articles := Articles{
		Article{
			Title:   "Test Title",
			Desc:    "Test Description",
			Content: "hello my love",
		},
	}

	fmt.Println("EndPoint Hit: All articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func main() {

	handleRequest()

}

func testPostArticles(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "Test Post endpoint worked")
}

//we see --> we get
func homePage(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "HomePage  one love")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
