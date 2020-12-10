package config

import (
	"os"
	"strings"
)

const (
	// Environment Variables
	EnvPort = "LISTEN_PORT"
)

var (
	Port string
)

//ReadEnv read the environment variables and store them in the configuration files
func ReadEnv() {
	var err error
	Port = os.Getenv(EnvPort)

}

// GetEnv returns the environment variable specified by key,
// or the default value specified by def if empty
func GetEnv(key string, def string) string {

	result := os.Getenv(key)
	if result == "" {
		return def
	}

	return strings.TrimSpace(result)
}
