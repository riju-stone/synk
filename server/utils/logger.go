package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func setupLogger() {
	log.SetOutput(os.Stdout)
	logLevel := log.WarnLevel

	if lvl, _ := os.LookupEnv("LOG_LEVEL"); lvl != "" {
		switch lvl {
		case "trace":
			logLevel = log.TraceLevel
		case "debug":
			logLevel = log.DebugLevel
		case "info":
			logLevel = log.InfoLevel
		case "warn":
			logLevel = log.WarnLevel
		case "error":
			logLevel = log.ErrorLevel
		default:
			logLevel = log.WarnLevel
		}
	}
	log.SetLevel(logLevel)
}

func InitLogger() {
	setupLogger()
}
