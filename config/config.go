package config

import (
	"net/http"
	"os"
	"strconv"
)

var (
	API_KEY       = getEnvWithDefault("API_ENV", "")
	CLIENT_ID     = getEnvWithDefault("CLIENT_ID", "1")
	CLIENT_SECRET = getEnvWithDefault("CLIENT_SECRET", "")
	SLEEP_TIME, _ = strconv.Atoi(getEnvWithDefault("SLEEP_TIME", "5"))
	SLACK_TOKEN   = getEnvWithDefault("SLACK_TOKEN", "")
	CLIENT        *http.Client
)

func Refresh() {
	initConfig()
}

func initConfig() {
	API_KEY = getEnvWithDefault("API_ENV", "")
	CLIENT_ID = getEnvWithDefault("CLIENT_ID", "1")
	CLIENT_SECRET = getEnvWithDefault("CLIENT_SECRET", "")
	SLEEP_TIME, _ = strconv.Atoi(getEnvWithDefault("SLEEP_TIME", "5"))
	SLACK_TOKEN = getEnvWithDefault("SLACK_TOKEN", "")
}

func getEnvWithDefault(key string, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}
