package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

type (
	User struct {
		gorm.Model
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
	}
)

var DB *gorm.DB
var err error
const(
	DNS = "root:admin@tcp(127.0.0.1:3306)/data?charset=utf8mb4&p"
)

func initialMigration(){
	DB, err = gorm.Open(mysql.Open(DNS), &&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error)
		panic("connot connect to Database..")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

}
