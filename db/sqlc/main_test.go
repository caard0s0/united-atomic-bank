package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DBDRIVER = "postgres"
	DBSOURCE = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
)

var TestQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(DBDRIVER, DBSOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	TestQueries = New(conn)

	os.Exit(m.Run())
}
