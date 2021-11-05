package main

import (
	"net/http"

	"logging"
	"routers"
)

func main() {
	logging.Init()

	router := routers.Route()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		logging.Save(err)
	}
}