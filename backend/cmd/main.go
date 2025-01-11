package main

import (
	"log"

	"github.com/FlamingoTP/cardz/helpers"
	"github.com/FlamingoTP/cardz/internal/api"
	"github.com/FlamingoTP/cardz/internal/database"
)

func main() {
	serverPort := helpers.GetEnvOrDefault("SERVER_PORT", "8080")

	databaseName := helpers.GetEnvOrDefault("DATABASE_NAME", "cardz")
	databaseAddress := helpers.GetEnvOrDefault("DATABASE_ADDRESS", "localhost:5432")
	databaseUsername := helpers.GetEnvOrDefault("DATABASE_USERNAME", "user")
	databasePassword := helpers.GetEnvOrDefault("DATABASE_PASSWORD", "s3cret")
	databaseArgs := helpers.GetEnvOrDefault("DATABASE_ARGS", "?sslmode=disable")

	dbCfg := database.Config{
		Address:      databaseAddress,
		Username:     databaseUsername,
		Password:     databasePassword,
		DatabaseName: databaseName,
		Args:         databaseArgs,
	}

	dbObj := &database.Database{Config: dbCfg}

	db := dbObj.Connect()

	defer db.Close()

	cfg := api.Config{
		Addr: ":" + serverPort,
	}

	app := &api.Application{Config: cfg}

	log.Fatal(app.Run(app.Mount()))
}
