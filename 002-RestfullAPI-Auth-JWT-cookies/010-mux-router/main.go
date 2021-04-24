package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Game struct {
		ID      string `json:"Id"`
		Title   string `json:"Title"`
		Company string `json:"company"`
	}

	GamesMark []Game
)

var Games []Game

func allGames(w http.ResponseWriter, re *http.Request) {
	games := GamesMark{
		Game{
			ID:      "1",
			Title:   "God of love",
			Company: "Test Description",
		},
		Game{
			ID:      "2",
			Title:   "Alone in the light",
			Company: "Test Description",
		},
	}

	fmt.Println("EndPoint Hit: All games Endpoint")
	json.NewEncoder(w).Encode(games)
}

func homePage(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "home page")
}

func returnSingleGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, game := range Games {
		if game.ID == key {
			json.NewEncoder(w).Encode(game)
		}
	}
}

func createNewGame(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var game Game
	json.Unmarshal(reqBody, &game)
	Games = append(Games, game)

	json.NewEncoder(w).Encode(game)
}

func deleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, game := range Games {
		if game.ID == id {
			Games = append(Games[:index], Games[index+1:]...)
		}
	}

}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homePage)
	r.HandleFunc("/games/{id}", allGames)
	r.HandleFunc("/game", createNewGame).Methods("POST")
	r.HandleFunc("/game/{id}", deleteGame).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9090", r))
}

func main() {
	handleRequests()
	println("serving at localhost: 9090")
}
