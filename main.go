package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tienvo76qng/simplepet/api"
	"github.com/tienvo76qng/simplepet/config"
	db "github.com/tienvo76qng/simplepet/db/sqlc"
)

func main() {

	config, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
