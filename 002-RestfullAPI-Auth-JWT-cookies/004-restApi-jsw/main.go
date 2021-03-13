package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("secret)")

func main() {
	/*
		    go get
		    github.com/dgrijalva/jwt-go

			to look for latest version
			go mod tidy
	*/
	fmt.Println("client service")
	println("----")
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	fmt.Fprint(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "fua"
	claims["expi"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)

	if err != nil {
		fmt.Errorf("error: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
