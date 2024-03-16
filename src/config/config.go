package config

import (
	"net/http"
	"os"
	"strconv"

	"golang.org/x/oauth2"
)

var API_KEY = getEnvWithDefault("API_ENV", "")
var CLIENT_ID = getEnvWithDefault("CLIENT_ID", "1")
var CLIENT_SECRET = getEnvWithDefault("CLIENT_SECRET", "")
var SLEEP_TIME, _ = strconv.Atoi(getEnvWithDefault("SLEEP_TIME", "5"))
var TOKEN *oauth2.Token
var SLACK_TOKEN = getEnvWithDefault("SLACK_TOKEN", "")
var CLIENT *http.Client

func getEnvWithDefault(key string, defaultValue string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return defaultValue
}
