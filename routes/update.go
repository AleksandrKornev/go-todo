package routes

import (
	"../schemas"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type resUpdate struct {
	Status string `json:"status"`
}

func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	var bodyText schemas.Todo
	err := json.NewDecoder(req.Body).Decode(&bodyText)
	if err != nil || bodyText.Id == 0 || len(bodyText.Text) == 0 {
		res.WriteHeader(400)
		var body resUpdate

		if bodyText.Id == 0 {
			body = resUpdate{"Not id"}
		} else if len(bodyText.Text) == 0 {
			body = resUpdate{"Not text"}
		}

		bodyJSON, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(res, string(bodyJSON))
		return
	}

	body := services.Update(bodyText)
	bodyJSON, _ := json.Marshal(body)

	res.WriteHeader(201)
	fmt.Fprintf(res, string(bodyJSON))
}
