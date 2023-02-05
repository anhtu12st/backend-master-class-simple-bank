package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	//config, err := util.LoadConfig("../..")
	//if err != nil {
	//	log.Fatal("cannot load config:", err)
	//}

	config := struct {
		DBDriver string
		DBSource string
	}{
		DBDriver: "postgres",
		DBSource: "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable",
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
