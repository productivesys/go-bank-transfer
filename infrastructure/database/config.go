package database

import (
	"os"
	"time"
)

type config struct {
	host     string
	database string
	port     string
	driver   string
	user     string
	password string

	connTimeout time.Duration
}

func newConfigMongoDB() *config {
	return &config{
		host:        os.Getenv("MONGODB_HOST"),
		database:    os.Getenv("MONGODB_DATABASE"),
		connTimeout: 60 * time.Second,
	}
}

func newConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		database: os.Getenv("POSTGRES_DATABASE"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}