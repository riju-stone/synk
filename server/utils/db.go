package utils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Connect to neon postgres database
// using environment variables for configuration
// and return a sql.DB instance
func ConnectDB() *pgx.Conn {
	// Load varabbles from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_CONNECTION_STRING")
	if dbUrl == "" {
		log.Fatal("DB_CONNECTION_STRING is not set")
	}
	log.Debug("DB_CONNECTION_STRING: ", dbUrl)

	// Connect to database
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.WithError(err).Fatal("Error connecting to database")
	}
	log.Info("Connected to database")

	// Check db health
	var version string
	err = conn.QueryRow(context.Background(), "select version()").Scan((&version))
	if err != nil {
		log.WithError(err).Fatal("Error checking database health")
	}
	log.Info("Database connection is healthy")

	return conn
}
