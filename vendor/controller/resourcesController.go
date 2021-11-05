package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
)

func ResourcesController(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	folder := vars["folder"]
	file := vars["file"]
	filePath := fmt.Sprintf("assets/%s/%s", folder, file)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		http.ServeFile(w, r, filePath)
	}
}