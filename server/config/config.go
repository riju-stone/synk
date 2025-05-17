package config

import (
	"os"
	"strconv"

	"github.com/riju-stone/synk/server/utils"
)

func GetLoggerConfig() *utils.LoggerConfig {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	logDir := os.Getenv("LOG_DIR")
	if logDir == "" {
		logDir = "./logs"
	}

	logFileName := os.Getenv("LOG_FILE_NAME")
	if logFileName == "" {
		logFileName = "app.log"
	}

	logMaxSize := os.Getenv("LOG_MAX_SIZE")
	if logMaxSize == "" {
		logMaxSize = "100"
	}

	maxSize, _ := strconv.Atoi(logMaxSize)

	return &utils.LoggerConfig{
		Level:       level,
		LogDir:      logDir,
		LogFileName: logFileName,
		LogMaxSize:  maxSize,
	}
}

func GetServerConfig() *utils.ServerConfig {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL is not set")
	}
	return &utils.ServerConfig{
		Host:        host,
		Port:        port,
		DatabaseURL: databaseURL,
	}
}
