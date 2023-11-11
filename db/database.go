package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"


)

var Bun *bun.DB

func CreateDatabase() (*sql.DB, error) {
	godotenv.Load()
	var (
		dbname = os.Getenv("DB_NAME")
		dbuser = os.Getenv("DB_USER")
		dbpass = os.Getenv("DB_PASS")
		dbhost = os.Getenv("DB_HOST")
		dbport = os.Getenv("DB_PORT")
		uri = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbuser, dbpass, dbname, dbhost, dbport)
	)

	db, err := sql.Open("postgres", uri)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Init() error {
	db, err := CreateDatabase()
	if err != nil {
		return err
	}

	Bun = bun.NewDB(db, pgdialect.New())
	return nil
}