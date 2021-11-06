package controller

import (
	"net/http"

	"logging"
)
func LoginController(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
	data["page"] = Page {
		Title: "ログインしてください",
	}
	view := View {
		View: "login/index",
		Writer: w,
	}
	err := view.Load(data)
	if err != nil {
	  logging.Save(err)
	}
}