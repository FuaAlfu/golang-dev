package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ClientRef struct {
	Client string `json:"client"`
	Data   struct {
		Name      string    `json: "name"`
		ID        int       `json: "id"`
		CreatedAT time.Time `json: "created_at"`
	}
}

func main() {

	//here
	c := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				// This is the TCP connect timeout in this instance.
				Timeout: 2500 * time.Millisecond,
			}).DialContext,
			TLSHandshakeTimeout: 2500 * time.Millisecond,
		},
	}
}
