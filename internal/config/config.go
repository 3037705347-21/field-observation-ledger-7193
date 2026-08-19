package config

import "os"

const DefaultPageSize = 50

type Settings struct {
	Address  string
	PageSize int
}

func (s Settings) WithDefaults() Settings {
	if s.Address == "" { s.Address = "127.0.0.1:18080" }
	if s.PageSize <= 0 { s.PageSize = DefaultPageSize }
	return s
}

func Load() Settings { return Settings{Address: os.Getenv("OBSERVE_ADDR")}.WithDefaults() }
