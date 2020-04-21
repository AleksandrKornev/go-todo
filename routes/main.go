package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type resRoot struct {
	Status bool `json:"status"`
}

func RootResponse(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.WriteHeader(200)

	body := resRoot{true}
	bodyJSON, _ := json.Marshal(body)

	fmt.Fprintf(res, string(bodyJSON))
}
