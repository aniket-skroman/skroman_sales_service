package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	// dbDriver = "postgres"
	// dbSource = "postgresql://postgres:support12@skroman-user.ckwveljlsuux.ap-south-1.rds.amazonaws.com:5432/skroman_sales_service"

	dbDriver = ""
	dbSource = ""
)

func makeConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	dbDriver = os.Getenv("DB_DRIVER")
	dbSource = os.Getenv("DB_SOURCE")

	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("db connection has been established..")
	return db
}

var DB_Instance = makeConnection()

func CloseDB(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
