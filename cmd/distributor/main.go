package main

import (
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/api"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/api/config"
	"log"
)

func main() {
	if err := api.Run(&config.Configuration{Addr: ":80"}); err != nil {
		log.Fatal(err)
	}
}
