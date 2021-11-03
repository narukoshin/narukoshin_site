package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	/*** MAIN ROUTES ***/
	router.HandleFunc("/", homeView)
	router.HandleFunc("/login", loginView)

	/*** RESOURCES ***/
	router.HandleFunc("/resources/{folderName}/{fileName}", resources)

	// Running server
	log.Print(http.ListenAndServe(":8000", router))
}

func homeView(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
		data["Site"] = map[string]string {
			"Title": "Hello, world",
		}
		t := template.Must(template.ParseFiles("templates/home.blade.html"))
		err := t.Execute(w, data)
		if err != nil {
			log.Print(err)
		}
}
func loginView(w http.ResponseWriter, r *http.Request){
	var userData map[string]string
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Print(err)
		}
		userName := r.Form["admin_username"][0]
		userPass := r.Form["admin_password"][0]

		userData = map[string]string {
			"username": userName,
			"password": userPass,
		}
	}
	data := make(map[string]interface{})
	data["Site"] = map[string]interface{} {
		"Title": "ログインをしてください",
		"userData": userData,
	}
	t := template.Must(template.ParseFiles("templates/login.blade.html"))
	err := t.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}
func resources(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	folderName := vars["folderName"]
	fileName := vars["fileName"]
	filePath := fmt.Sprintf("assets/%s/%s", folderName, fileName)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		http.ServeFile(w, r, filePath)
	} else {
		log.Print(err)
	}
}