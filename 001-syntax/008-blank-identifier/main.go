package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	user struct {
		ID   int
		Name string
	}

	updateState struct {
		Modified int
		Duration float64
		Success  bool
		Message  string
	}
)

func main() {
	/*
		database practice
	*/
	u := user{
		ID:   1578,
		Name: "john",
	}

	//here we used a blank identifier: _
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	//Display the update was Successful
	fmt.Println("update user record for id ", u.ID)
}

//update-user updates the specified  user docment
func updateUser(u *user) (*updateState, error) {
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "hi, everything is awesome"}`

	//unmarshal the json docment into a value of
	var us updateState
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	//check the update status to verify the update
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
