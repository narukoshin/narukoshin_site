package controller

import "net/http"

func HomeController(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
	data["page"] = Page {
		Title: "hello, world",
	}

	view := View {
		View: "home/index",
		Writer: w,
	}
	view.Load(data)
}