package routes

import (
	"../schemas"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type resCreateErr struct {
	Status string `json:"status"`
}

type resCreate struct {
	Id int `json:"id"`
}

func Create(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	var bodyText schemas.Todo
	err := json.NewDecoder(req.Body).Decode(&bodyText)
	if err != nil || len(bodyText.Text) == 0 {
		res.WriteHeader(400)
		body := resCreateErr{"Not text"}
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(res, string(bodyJSON))
		return
	}

	body := resCreate{services.Create(bodyText)}
	bodyJSON, _ := json.Marshal(body)

	res.WriteHeader(201)
	fmt.Fprintf(res, string(bodyJSON))
}
