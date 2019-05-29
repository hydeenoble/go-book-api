package config

import (
	"log"
	"os"
	"strconv"
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
	postgresUsername := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresDbName := os.Getenv("POSTGRES_DB")
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
			Username: postgresUsername,
			Password: postgresPassword,
			Host:     postgresHost,
			DbName:   postgresDbName,
			Port:     postgresPort,
		},
	}
}
