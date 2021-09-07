package main

import (
	"os"

	"google.golang.org/grpc"
)

func init() {
	grpc.Dial(os.Getenv("ENDPOINT"), grpc.WithInsecure())
}
