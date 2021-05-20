package main

import (
	"net/http"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
	}
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

}
