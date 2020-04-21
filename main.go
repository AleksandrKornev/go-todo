package main

import (
	"./initialize"
	"./routes"
	"net/http"
)

func main() {
	initialize.Init()

	http.HandleFunc("/", routes.RootResponse)
	http.HandleFunc("/get", routes.Get)
	http.HandleFunc("/create", routes.Create)
	http.HandleFunc("/update", routes.Update)
	http.HandleFunc("/remove", routes.Remove)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
