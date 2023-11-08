package querytest

import (
	"database/sql"
	"log"
	"os"
	"testing"

	sqlc_lib "github.com/aniket-skroman/skroman_sales_service.git/sqlc_lib"
	_ "github.com/lib/pq"
)

var (
	testQueries *sqlc_lib.Queries
	testDB      *sql.DB
	db_driver   = "postgres"
	db_source   = "postgresql://postgres:support12@skroman-user.ckwveljlsuux.ap-south-1.rds.amazonaws.com:5432/skroman_sales_service"
)

func TestMain(t *testing.M) {
	var err error
	testDB, err = sql.Open(db_driver, db_source)

	if err != nil {
		log.Fatal("failed to connect db : ", err)
	}

	testQueries = sqlc_lib.New(testDB)
	os.Exit(t.Run())
}
