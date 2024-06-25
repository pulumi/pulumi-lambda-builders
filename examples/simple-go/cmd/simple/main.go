package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event map[string]string) (*string, error) {
	message := "Hello World!"
	return &message, nil
}

func main() {
	lambda.Start(handleRequest)
}
