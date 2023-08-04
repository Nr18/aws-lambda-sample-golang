package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func Handler(ctx context.Context, event MyEvent) {
	log.Printf("Welcome %s", event.Name)
}

func main() {
	lambda.Start(Handler)
}