package main

import (
	"github.com/rubenwo/DistributedTranscoding/pkg/cli"
	"log"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}
