package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Key string `json:"key"` // Statuspage API key
}

func ParseConfig() (c *Config) {
	// Read config
	const path = "config.json"
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Parse config
	c = &Config{}
	err = json.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}

	// Validate config
	if len(c.Key) == 0 {
		panic("missing Statuspage API key in config.json")
	}

	return c
}
