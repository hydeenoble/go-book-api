package config

import (
	"log"
	"os"
	"strconv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	API   *APIConfig
	Mongo *MongoConfig
}

type APIConfig struct {
	Name string
	Port int
}

type MongoConfig struct {
	Username string
	Password string
	Host     string
	DbName   string
	Port     int
}

func GetConfig() *Config {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))
	mongoPort, mongoPortErr := strconv.Atoi(os.Getenv("MONGO_PORT"))

	if portErr != nil {
		log.Fatal(portErr)
	}

	if mongoPortErr != nil {
		log.Fatal(mongoPortErr)
	}

	return &Config{
		API: &APIConfig{
			Name: "Sample Book API In Golang",
			Port: port,
		},
		Mongo: &MongoConfig{
			Username: os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Host:     os.Getenv("MONGO_HOST"),
			DbName:   os.Getenv("MONGO_DB"),
			Port:     mongoPort,
		},
	}
}
