package controller

import (
	"net/http"
	"fmt"
	"os"
	
	"github.com/gorilla/mux"
	"narukoshin.me/logging"
)

func ResourcesController(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	folder := vars["folder"]
	file := vars["file"]
	filePath := fmt.Sprintf("assets/%s/%s", folder, file)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		http.ServeFile(w, r, filePath)
	} else {
		logging.Save(err)
	}
}