// Package database contains the logic to open connections to a sqlite database.
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "quotations.db")
	if err != nil {
		log.Println("error creating database connection: ", err.Error())
		return nil, err
	}

	err = runMigrations(db)
	if err != nil {
		log.Println("error when running database migrations: ", err.Error())
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS [quotation] (%s, %s, %s, %s);`,
		"code TEXT NOT NULL",
		"code_in TEXT NOT NULL",
		"bid REAL NOT NULL",
		"PRIMARY KEY(code, code_in)",
	)
	_, err := db.Exec(query)

	return err
}
