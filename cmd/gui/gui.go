package main

import (
	"github.com/rubenwo/DistributedTranscoding/pkg/client"
	"github.com/rubenwo/DistributedTranscoding/pkg/gui"
)

func main() {
	gui.Run(client.New())
}
