package main

import (
	"net/http"
	"html/template"
)

func ErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	template.Must(template.ParseGlob("template/500Error.html"))
	sere.ExecuteTemplate(w, "template/500Error.html", nil)
}