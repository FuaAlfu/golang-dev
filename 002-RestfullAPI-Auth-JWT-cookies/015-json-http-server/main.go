package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	BirthYear int    `json:"birth_year"`
	Email     string
}

func Marshal(v interface{}) ([]byte, error)
func Unmarshal(data []byte, v interface{}) error

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) error {
	return respondWithJSON(w, code, map[string]string{"error": msg})
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type responseBody struct {
		Token string `json:"token"`
	}

	user, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 500, "couldn't read request")
		return
	}
	params := requestBody{}
	err = json.Unmarshal(user, &params)
	if err != nil {
		respondWithError(w, 500, "couldn't unmarshal parameters")
		return
	}

	// do stuff with username and password

	respondWithJSON(w, 200, responseBody{
		Token: "example-auth-token",
	})
}

func main() {
	fmt.Println("constructing all local users")
	u1, _ := json.Marshal(User{
		FirstName: "jhon",
		BirthYear: 2003,
		Email:     "example@gmail.com",
	})
	fmt.Println(string(u1))
	println("---")
	u2, _ := json.Marshal(User{
		FirstName: "doe",
		BirthYear: 1994,
		Email:     "example2@gmail.com",
	})
	fmt.Println(string(u2))
	println("---")
	user1 := User{}
	err := json.Unmarshal(u1, &user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user1)
	println("---")
	user2 := User{}
	err = json.Unmarshal(u2, &user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user2)
}
