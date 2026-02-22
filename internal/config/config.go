package config

import (
	"log"
	"os"
)

type Config struct {
	Addr   string
	DBDSN  string
	DBCERT string
	Env    string
}

func Load() Config {
	addr := os.Getenv("APP_ADDR")
	dbdsn := os.Getenv("DB_DSN")
	dbcert := os.Getenv("DB_CERT")
	env := os.Getenv("APP_ENV")

	if addr == "" {
		log.Fatal("Server Address Enviroment failed to load!")
	}

	return Config{
		Addr:   addr,
		DBDSN:  dbdsn,
		DBCERT: dbcert,
		Env:    env,
	}
}
