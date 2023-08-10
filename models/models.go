package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "wordcards.db")
	if err != nil {
		return err
	}
	return nil
}

func CreateTables() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS word_cards (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			front VARCHAR(255),
			back VARCHAR(255)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}
