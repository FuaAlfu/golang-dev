package main

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//our key for jwt to sign up
var jwtkey = []byte("secret")

//create a simple data
var user = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

type (
	Credentials struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}

	Claims struct {
		Username string `json: "username"`
		jwt.StandardClaims
	}
)

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]
	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//todo
}

func Home(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
