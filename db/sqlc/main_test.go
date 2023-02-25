package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Defined as global variable
var testQueries *Queries

var db *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// Main entry point of all unit test by convention
func TestMain(m *testing.M) {
	var error error

	db, error = sql.Open(dbDriver, dbSource)

	if error != nil {
		log.Fatal("Cannot connect to the database.")
	}

	testQueries = New(db)

	os.Exit(m.Run())
}
