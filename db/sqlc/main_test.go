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
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(DBDRIVER, DBSOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	TestQueries = New(testDB)

	os.Exit(m.Run())
}
