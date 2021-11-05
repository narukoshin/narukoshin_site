package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type View struct {
	View string
}

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

// Home view
func homeView(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
	data["Site"] = map[string]string {
		"Title": "Hello, world",
	}

	// Loading a view
	view := View {View: "home"}
	view.load(w, data)
}

// Login view
func loginView(w http.ResponseWriter, r *http.Request){
	var userData map[string]string
	if r.Method == "POST" {
		// parsing the form
		err := r.ParseForm()
		if err != nil {
			log.Print(err)
		}

		// getting the username and password
		userName := r.Form["admin_username"][0]
		userPass := r.Form["admin_password"][0]

		// setting the username and password in userData
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

	// Loading a view
	view := View {View: "login"}
	view.load(w, data)
}

// Loading resources
func resources(w http.ResponseWriter, r *http.Request){
	// parsing the parameters
	vars := mux.Vars(r)

	// getting folder name from the parameters
	folderName := vars["folderName"]

	// getting file name from the parameters
	fileName := vars["fileName"]

	// building a path
	filePath := fmt.Sprintf("assets/%s/%s", folderName, fileName)

	// checking if file exists
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		// serving the file to the client
		http.ServeFile(w, r, filePath)
	} else {
		// if file doesn't exists, then print the error
		log.Print(err)
	}
}

func (v *View) load(w http.ResponseWriter, data interface{}) {
	// checking if the view file exists
	fileLocation := "templates/" + v.View + ".blade.html"
	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		// if the file exists
		t := template.Must(template.ParseFiles(fileLocation))
		// Executing template
		err := t.Execute(w, data)
		if err != nil {
			log.Print(err)
		}
	} else {
		// if template file doesn't exists
		w.Write([]byte("Template file doesn't not exists"))
	}
}