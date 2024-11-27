package helpers

import "os"

func GetEnvOrDefault(envVar, defaultVal string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}
	return defaultVal
}
