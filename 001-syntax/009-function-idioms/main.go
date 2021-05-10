package main

import(
	"fmt"
	"encoding/json"
)

type(
	user struct{
		ID int
		Name string
		Password string
	}
)

func main() {
	//retrive the user profile
	u, err := retriveUser("mua")
	if err != nil{
		fmt.Println(err)
	}

	//display the user profile
	fmt.Println("%+v\n", *u)
}

//return multaple values out of function , return value and the other state in the stame time
func retriveUser(name string)(&user, error){
	//make a call to get the user in json respone..
	r, err := getUser(name)
	if err != nil{
		return nil, err
	}
	/*
	unmarshal the json document into a value 
	of the user struct type.
	*/
	var u user
	if err := json.Unmarshal([]byte(r), &u); err != nil{
	return nil, err
}
	return &u, nil
}

//getuser simulates a webcall that returns a json docment for the specified user.
func getUser(name string)(string,error){
	response := `{"id": 1234, "name":"mua"}`
	return response, nil
}