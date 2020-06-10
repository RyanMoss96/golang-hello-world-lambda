package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type MyResponse struct {
	Message string `json:"message:"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return MyResponse{Message: fmt.Sprintf("Hello World")}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
