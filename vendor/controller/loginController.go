package controller

import "net/http"

func LoginController(w http.ResponseWriter, r *http.Request){
	data := make(map[string]interface{})
	data["page"] = Page {
		Title: "ログインしてください",
	}
	view := View {
		View: "login/index",
		Writer: w,
	}
	view.Load(data)
}