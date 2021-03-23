package main

import (
	"context"
	"github.com/rubenwo/DistributedTranscoding/pkg/transcoder"
	"log"
)

func main() {
	client := transcoder.NewClient("localhost:8080", 4)
	ctx := context.TODO()
	log.Println("joining cluster")
	if err := client.JoinCluster(ctx); err != nil {
		log.Fatal(err)
	}
}
