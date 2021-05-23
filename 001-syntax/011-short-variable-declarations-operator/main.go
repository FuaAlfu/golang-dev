package main

import (
	"fmt"
)

type (
	user struct {
		id   int
		name string
	}

	customer struct {
		id   int
		name string
	}
)

func main() {
	var err error

	/*
		the short variable declaration operation will
		declare u and redeclare err
	*/
	u, err := getUser()
	if err != nil {
		return
	}

	fmt.Println(u)
	println("---")

	c, err := getCustomer(u)
	if err != nil {
		return
	}

	fmt.Println(c)
}

//return a pointer of type user
func getUser() (*user, error) {
	return &user{65499, "john"}, nil
}

func getCustomer(u *user) (*customer, error) {
	return &customer{65499, "john"}, nil
}
