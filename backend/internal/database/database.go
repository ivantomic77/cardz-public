package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Config Config
}

type Config struct {
	Address      string
	Username     string
	Password     string
	DatabaseName string
	Args         string
}

func (db *Database) Connect() *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s%s", db.Config.Username, db.Config.Password, db.Config.Address, db.Config.DatabaseName, db.Config.Args)

	dbInstance, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = dbInstance.Ping(); err != nil {
		_ = dbInstance.Close()
		log.Fatal(err)
	}

	log.Printf("Server successfully connected to database '%s' on address '%s'", db.Config.DatabaseName, db.Config.Address)

	return dbInstance
}
