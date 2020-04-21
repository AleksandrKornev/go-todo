package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DB() *sql.DB {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err.Error())
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
