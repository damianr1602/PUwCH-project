package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config represents application connector credentials and db name
type Config struct {
	Server   string
	Database string
}

// Read and parse the config file
func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
