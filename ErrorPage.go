package main

import (
	"net/http"
	// "html/template"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	http.ServeFile(w, r, "template/500Error.html")
	// template.Must(template.ParseGlob("template/500Error.html"))
	// sere.ExecuteTemplate(w, "template/500Error.html", nil)
}