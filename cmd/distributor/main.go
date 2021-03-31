package main

import (
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/api"
	"github.com/rubenwo/DistributedTranscoding/pkg/distributor/api/config"
	"log"
)

func main() {
	if err := api.Run(&config.Configuration{ApiAddr: ":80", AltApiAddr: ":81", ClusterAddr: ":8080"}); err != nil {
		log.Fatal(err)
	}
}
