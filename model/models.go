package model

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


