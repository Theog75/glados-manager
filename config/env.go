package config

import (
	"os"
	"strings"
)

const (
	// Environment Variables
	EnvPort           = "LISTEN_PORT"
	EnvKubeConfigPath = "KUBE_CONFIG_PATH"
)

var (
	Port           string
	KubeConfigPath string
)

//ReadEnv read the environment variables and store them in the configuration files
func ReadEnv() {
	Port = os.Getenv(EnvPort)
	KubeConfigPath = GetEnv(EnvKubeConfigPath, "")
	if KubeConfigPath == "" {

		// fmt.Println(EnvKubeConfigPath, "is not set.  Using default.")

		KubeConfigPath = "/ocpcluster-kube/ocpcluster-config-3.11"

	}

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
