package main

import (
	"plugin"

	// This line causes the issue.
	_ "google.golang.org/grpc"
)

func main() {
	plugin.Open("plugin.so")
}
