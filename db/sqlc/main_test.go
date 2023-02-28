package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/igua95/simplebank/util"
	_ "github.com/lib/pq"
)

// Defined as global variable
var testQueries *Queries

var db *sql.DB

// Main entry point of all unit test by convention
func TestMain(m *testing.M) {
	var error error
	var config util.Config

	config, error = util.LoadConfig("../../.")
	if error != nil {
		log.Fatal("Cannot read config.", error)
	}

	db, error = sql.Open(config.DBDriver, config.DBSource)

	if error != nil {
		log.Fatal("Cannot connect to the database.")
	}

	testQueries = New(db)

	os.Exit(m.Run())
}
