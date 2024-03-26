package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var config = map[string]string{
	// general
	"SERVICE_NAME":        "to_do_list",
	"SERVICE_VERSION":     "0.0.0",
	"ENV":                 "local",
	"LOG_LEVEL":           "debug",
	"ENABLE_DD_PROFILING": "false",
	"ENABLE_GO_PROFILING": "false",

	// http server
	"HTTP_PORT":                        "9000",
	"HTTP_SERVER_READ_TIMEOUT_MILLIS":  "60000",
	"HTTP_SERVER_WRITE_TIMEOUT_MILLIS": "60000",
	"HTTP_AUTH_TOKEN":                  "altFmZK2YbjyYU0wXLkqJB0KZ6fSt5PIQMoCCkS67a4iobhyKd45M3a8riMGIC0g",

	// fake http server
	"FAKE_HTTP_PORT": "9001",

	// postgresql
	"DB_HOST":                "localhost",
	"DB_PORT":                "5432",
	"DB_USER":                "to-do-list-user",
	"DB_PASSWORD":            "to-do-list-password",
	"DB_NAME":                "to-do-list",
	"DB_POOL_SIZE":           "5",
	"DB_CONN_MAX_TTL_MILLIS": "1800000",
	"DB_TIMEOUT_SECONDS":     "2",
	"DB_LOCK_TIMEOUT_MILLIS": "800",
}

func GetString(k string) string {
	v := os.Getenv(k)
	if v == "" {
		return config[k]
	}

	return v
}

// GetInt value of a given env var
func GetInt(k string) int {
	v := GetString(k)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}

	return i
}

// GetDuration value of a given env var
func GetDuration(k string) time.Duration {
	return time.Duration(GetInt(k)) * time.Millisecond
}

// GetBool value of a given env var
func GetBool(k string) bool {
	v := GetString(k)
	return strings.ToLower(v) == "true"
}

// Set config for test purposes
func Set(k, v string) {
	config[k] = v
}