package initialize

import (
	"../helpers"
)

func Init() {
	db := helpers.DB()
	defer db.Close()

	var isHave bool

	_, err := db.Exec("SELECT * FROM todo")
	if err != nil {
		isHave = false
	}

	if isHave == false {
		fillTable()
	}
}

func fillTable() {
	db := helpers.DB()
	defer db.Close()

	db.Exec(`
	CREATE TABLE todo (
		id   INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, 
    text text,
    date integer
	)
	`)
}
