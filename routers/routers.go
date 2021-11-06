package routers

import (
	"github.com/gorilla/mux"
	"narukoshin.me/controller"
)

func Route() *mux.Router {
	/*** Starting the Mux Router ***/
	route := mux.NewRouter()

	/*** Handling the routes ***/
	route.HandleFunc("/", controller.HomeController)
	route.HandleFunc("/login", controller.LoginController)
	route.HandleFunc("/resources/{folder}/{file}", controller.ResourcesController)

	/*** Returning mux.Router ***/
	return route
}