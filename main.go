package main

import (
	"database/sql"
	"log"

	"github.com/igua95/simplebank/api"
	db "github.com/igua95/simplebank/db/sqlc"
	"github.com/igua95/simplebank/util"
	_ "github.com/lib/pq" // to talk to de DB
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot read config.", err)
	}

	connection, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the database.", err)
	}

	store := db.NewStore(connection)

	server := api.NewService(store)

	err = server.Start(config.ServerAddres)

	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}
