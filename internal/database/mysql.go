package database

import (
	"database/sql"
	"os"

	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQL(dsn string) (*sql.DB, error) {
	// Register custom TLS config with RDS CA
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile("/etc/ssl/certs/aws-ad-bundle.pem")
	if err != nil {
		log.Fatalf("Failed to read cert: %v", err)
	}
	rootCertPool.AppendCertsFromPEM(pem)
	mysql.RegisterTLSConfig("custom", &tls.Config{RootCAs: rootCertPool})

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping: %v", err)
	}

	var v string
	if err := db.QueryRow("SELECT VERSION()").Scan(&v); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	fmt.Println(v)
	return db, nil

}
