package main

import (
	"fmt"

)

type (
	HelloRequest struct {
		Name string
	}

	HelloResponse struct {
		Greet string
	}

	WelcomeServiceClient interface {
		Hello(*HelloRequest) (*HelloResponse, error)
	}

	WelcomeServiceServer interface {
		Hello(*HelloRequest) (*HelloResponse, error)
	}
)

func main() {
	fmt.Println("hi")
}
