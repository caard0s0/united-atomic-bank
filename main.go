package main

import (
	"atomic-bank/api"
	db "atomic-bank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	DBDRIVER      = "postgres"
	DBSOURCE      = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
	SERVERADDRESS = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(DBDRIVER, DBSOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(SERVERADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
