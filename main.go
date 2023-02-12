package main

import (
	"backend-master-class-simple-bank/api"
	db "backend-master-class-simple-bank/db/sqlc"
	"backend-master-class-simple-bank/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

}
