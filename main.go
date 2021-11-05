package main

import (
	"log"
	"net/http"
	"routers"
)

func main() {
	router := routers.Route()
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Print(err)
	}
}