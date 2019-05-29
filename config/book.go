package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	API *APIConfig
}

type APIConfig struct {
	Name string
	Port int
}

func GetConfig() *Config {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))
	if portErr != nil {
		log.Fatal(portErr)
	}
	return &Config{
		API: &APIConfig{
			Name: "Sample Book API In Golang",
			Port: port,
		},
	}
}
