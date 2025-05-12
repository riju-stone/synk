package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var StunServers = []string{
	"stun:stun.l.google.com:19302",
	"stun:stun.l.google.com:5349",
	"stun:stun1.l.google.com:3478",
}

type ServerConfig struct {
	logLevel    string
	domain      string
	port        int
	stunServers *[]string
}

func LoadConfig() (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	logLevel := os.Getenv("LOG_LEVEL")
	domain := os.Getenv("DOMAIN")
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	stunServers := &StunServers

	return &ServerConfig{
		logLevel:    logLevel,
		domain:      domain,
		port:        portInt,
		stunServers: stunServers,
	}, nil
}
