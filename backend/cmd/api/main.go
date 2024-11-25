package main

import (
	"log"
	"os"

	"github.com/FlamingoTP/cardz/internal/api"
)

func main() {
	addr := getEnvOrDefault("SERVER_PORT", "8080")

	cfg := api.Config{
		Addr: ":" + addr,
	}

	app := &api.Application{Config: cfg}

	log.Fatal(app.Run(app.Mount()))
}

func getEnvOrDefault(envVar, defaultVal string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}
	return defaultVal
}
