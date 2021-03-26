package main

import (
	"net/http"
	"sync"

	_ "github.com/dgrijalva/jwt-go"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type authHandler struct {
	key []byte
}

var mySecretKey = []byte("")

func unauthorized(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}

func main() {
	r := http.NewServeMux()
	user := &userHandler{
		save: &datastore{
			m: map[string]user{
				"1": user{ID: "1", Name: "fua"},
			},
			RWMutex: &sync.RWMutex{},
		},
		key: mySecretKey,
	}
	r.Handle("/users", user)
	r.Handle("/users", user)

	http.ListenAndServe(":9090", r)
}
