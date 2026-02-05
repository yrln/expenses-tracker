package config

import (
	"log"
	"os"
)

type Config struct {
	Addr  string
	DBDSN string
}

func Load() Config {
	addr := os.Getenv("APP_ADDR")
	dbdsn := os.Getenv("DB_DSN")

	if addr == "" {
		log.Fatal("Server Address Enviroment failed to load!")
	}

	return Config{
		Addr:  addr,
		DBDSN: dbdsn,
	}
}
