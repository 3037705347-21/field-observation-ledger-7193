package config

import "os"

type Settings struct {
	Address string
}

func Load() Settings {
	address := os.Getenv("OBSERVE_ADDR")
	if address == "" {
		address = "127.0.0.1:18080"
	}
	return Settings{Address: address}
}
