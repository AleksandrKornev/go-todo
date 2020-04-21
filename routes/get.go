package routes

import (
	"../schemas"
	"../services"
	"encoding/json"
	"fmt"
	"net/http"
)

type resGet struct {
	Status string `json:"status"`
	items  []schemas.Todo
}

func Get(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

	body := services.Get()
	bodyJSON, _ := json.Marshal(body)

	res.WriteHeader(201)
	fmt.Fprintf(res, string(bodyJSON))
}
