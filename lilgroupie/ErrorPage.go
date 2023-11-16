package lilGroupie

import (
	"net/http"
	// "html/template"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, err int) {
	if err == 400 {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "template/400Error.html")
	} else if err == 404 {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")
	} else if err == 405 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	} else if err == 500 {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "template/500Error.html")
	}

	// template.Must(template.ParseGlob("template/500Error.html"))
	// sere.ExecuteTemplate(w, "template/500Error.html", nil)
}
