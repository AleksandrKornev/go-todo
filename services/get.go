package services

import (
	"../helpers"
	"../schemas"
)

func Get() []schemas.Todo {
	var items []schemas.Todo
	query := "SELECT * FROM todo"

	db := helpers.DB()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		row := schemas.Todo{}
		err := rows.Scan(&row.Id, &row.Text, &row.Date)
		if err != nil {
			continue
		}

		items = append(items, row)
	}

	return items
}
