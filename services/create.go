package services

import (
	"../helpers"
	"../schemas"
	"strconv"
	"time"
)

func Create(todo schemas.Todo) int {
	date := time.Now().Unix()
	dateStr := strconv.Itoa(int(date))
	query := `
	INSERT INTO todo 
	(text, date) 
	values ('` + todo.Text + "'," + dateStr + ")"

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
