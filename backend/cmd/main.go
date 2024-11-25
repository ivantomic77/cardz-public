package main

import (
	"log"

	"github.com/FlamingoTP/cardz/helpers"
	"github.com/FlamingoTP/cardz/internal/api"
)

func main() {
	addr := helpers.GetEnvOrDefault("SERVER_PORT", "8080")

	cfg := api.Config{
		Addr: ":" + addr,
	}

	app := &api.Application{Config: cfg}

	log.Fatal(app.Run(app.Mount()))
}
