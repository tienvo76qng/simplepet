package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/tienvo76qng/simplepet/config"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := config.LoadConfig("../../config")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
