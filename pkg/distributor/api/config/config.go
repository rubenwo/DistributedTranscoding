package config

import "fmt"

type Configuration struct {
	// ApiAddr is the address where clients can submit a job on using REST/JSON
	ApiAddr string

	// AltApiAddr is the address where clients can submit a job on using gRPC
	AltApiAddr string

	// ClusterAddr is the address where transcoders can join on using gRPC
	ClusterAddr string
}

func (cfg *Configuration) Validate() error {
	if cfg.ApiAddr == "" {
		return fmt.Errorf("ApiAddr can't ben empty")
	}
	if cfg.AltApiAddr == "" {
		return fmt.Errorf("AltApiAddr can't ben empty")
	}
	if cfg.ClusterAddr == "" {
		return fmt.Errorf("ClusterAddr can't ben empty")
	}

	if cfg.ApiAddr == cfg.AltApiAddr {
		return fmt.Errorf("ApiAddr can't be the same as AltApiAddr")
	}
	if cfg.ApiAddr == cfg.ClusterAddr {
		return fmt.Errorf("ApiAddr can't be the same as ClusterAddr")
	}
	if cfg.ClusterAddr == cfg.AltApiAddr {
		return fmt.Errorf("ClusterAddr can't be the same as AltApiAddr")
	}

	return nil
}
