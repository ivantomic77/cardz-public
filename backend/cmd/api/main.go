package main

import (
	"log"
	"os"
)

func main() {
	addr := getEnvOrDefault("SERVER_PORT", "8080")

	cfg := config{
		addr: ":" + addr,
	}

	app := &application{config: cfg}

	log.Fatal(app.run(app.mount()))
}

func getEnvOrDefault(envVar, defaultVal string) string {
	if value := os.Getenv(envVar); value != "" {
		return value
	}
	return defaultVal
}
