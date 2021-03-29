package main

import (
	"fmt"
)

type (
	user struct {
		id    int
		email string
	}
	account struct{ users []user }
)

func main() {
	fua := user{id: 1, email: "fua@fua"}
	mua := user{id: 2, email: "mua@mua"}
	nata := user{id: 3, email: "nata@nata"}

	accounts := account{
		users: []user{fua, mua, nata},
	}

	fmt.Println(fua, mua, nata)
	println("---")
	fmt.Println(accounts.users)
}
