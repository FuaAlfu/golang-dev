package main

import (
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqLite"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func initialMigaration() {
	db, err := gorm.Open("sqlite3", "myData.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect..")
	}
	db.Close()
	db.AutoMigrate(&User{})
}

func homePage(w http.ResponseWriter, re *http.Request) {
	fmt.Fprintf(w, "test..")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
