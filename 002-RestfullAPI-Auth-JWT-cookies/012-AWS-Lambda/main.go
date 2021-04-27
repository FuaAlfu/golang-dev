package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type (
	Request struct {
		ID    float64 `json:"id"`
		Value string  `json:"value"`
	}

	Response struct {
		Message string `json:"message"`
	}
)

func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
