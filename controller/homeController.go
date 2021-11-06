package controller

import (
	"net/http"

	"logging"
)

func HomeController(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
	data["page"] = Page {
		Title: "hello, world",
	}

	view := View {
		View: "home/index",
		Writer: w,
	}
	err := view.Load(data)
	if err != nil {
		logging.Save(err)
	}
}