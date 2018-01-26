package main

import (
	"os"
)

// Config struct for cExporter
type Config struct {
	httpHost string
	httpPort string
}

// GetConfig - Return config for service
func GetConfig() Config {
	var config Config

	if config.httpHost = os.Getenv("HTTP_HOST"); config.httpHost == "" {
		config.httpHost = "127.0.0.1"
	}

	if config.httpPort = os.Getenv("HTTP_PORT"); config.httpPort == "" {
		config.httpPort = "9167"
	}

	return config
}
