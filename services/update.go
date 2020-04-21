package services

import (
	"../helpers"
	"../schemas"
	"strconv"
)

func Update(todo schemas.Todo) int {
	query := `
	UPDATE todo
	SET text = '` + todo.Text + `'
	WHERE id = ` + strconv.Itoa(int(todo.Id))

	db := helpers.DB()
	defer db.Close()

	row, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	messageID, err := row.LastInsertId()
	if err != nil {
		panic(err)
	}

	return int(messageID)
}
