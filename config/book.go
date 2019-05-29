package config

import (
	"log"
	"os"
	"strconv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	API      *APIConfig
	Postgres *PostgresConfig
}

type APIConfig struct {
	Name string
	Port int
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	DbName   string
	Port     int
}

func GetConfig() *Config {
	port, portErr := strconv.Atoi(os.Getenv("PORT"))
	postgresPort, postgresPortErr := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	if portErr != nil {
		log.Fatal(portErr)
	}

	if postgresPortErr != nil {
		log.Fatal(postgresPortErr)
	}

	return &Config{
		API: &APIConfig{
			Name: "Sample Book API In Golang",
			Port: port,
		},
		Postgres: &PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			DbName:   os.Getenv("POSTGRES_DB"),
			Port:     postgresPort,
		},
	}
}
