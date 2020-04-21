package services

import (
	"../helpers"
	"../schemas"
	"strconv"
)

func Remove(todo schemas.Todo) {
	query := `
	DELETE FROM todo
	WHERE id = ` + strconv.Itoa(int(todo.Id))

	db := helpers.DB()
	defer db.Close()

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
