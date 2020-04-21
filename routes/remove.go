package routes

import (
	"../schemas"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type resRemove struct {
	Status string `json:"status"`
}

func Remove(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	var bodyText schemas.Todo
	err := json.NewDecoder(req.Body).Decode(&bodyText)
	if err != nil || bodyText.Id == 0 {
		res.WriteHeader(400)
		body := resRemove{"Not id"}

		bodyJSON, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(res, string(bodyJSON))
		return
	}

	services.Remove(bodyText)
	body := resRemove{"Success"}
	bodyJSON, _ := json.Marshal(body)

	res.WriteHeader(201)
	fmt.Fprintf(res, string(bodyJSON))
}
