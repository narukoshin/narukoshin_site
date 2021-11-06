package main

import (
	"net/http"

	"narukoshin.me/logging"
	"narukoshin.me/routers"

	"github.com/nanmu42/gzip"
)

func main() {
	logging.Init()

	router := routers.Route()

	err := http.ListenAndServe(":8000", gzip.DefaultHandler().WrapHandler(router))
	if err != nil {
		logging.Save(err)
	}
}