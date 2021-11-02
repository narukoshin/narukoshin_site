package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["Site"] = map[string]string {
			"Title": "Hello, world",
		}
		t := template.Must(template.ParseFiles("templates/home.blade.html"))
		t.Execute(w, data)
	})
	router.HandleFunc("/assets/video/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		location := "assets/video/" + name
		if _, err := os.Stat(location); !os.IsNotExist(err) {
			http.ServeFile(w, r, location)
		}
	})
	router.HandleFunc("/vendor/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		location := "vendor/" + name
		if _, err := os.Stat(location); !os.IsNotExist(err) {
			http.ServeFile(w, r, location)
		} 
	})

	// giving access to the stylesheet folder
	router.HandleFunc("/assets/css/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		location := "assets/css/" + name
		if _, err := os.Stat(location); !os.IsNotExist(err) {
			http.ServeFile(w, r, location)
		}
	})
	http.ListenAndServe(":8000", router)
}