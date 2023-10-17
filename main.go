package main

import (
	// "bufio"
	"fmt"
	"html/template"
	// "io"
	"log"
	"net/http"
	"net/url"
	// "os"
	// "strings"
)

var sere *template.Template

func init() {
	sere = template.Must(template.ParseGlob("template/index.html"))
}

func main() {
	http.HandleFunc("/", Index)
	http.Handle("/template/css/", http.StripPrefix("/template/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/groupie-tracker", processor)
	// http.HandleFunc("/download", DownLoad)
	u, err := url.Parse("http://localhost:2003")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Listening and serving on: ")
	fmt.Printf("%+v", u)
	fmt.Println()
	log.Fatal(http.ListenAndServe(":2003", nil))

}


func Index(w http.ResponseWriter, r *http.Request) {
	FromJasonToTxt(w)
	if r.Method == "GET" {
		if r.URL.Path == "/" {
			sere.ExecuteTemplate(w, "index.html", nil)
			return
		} else {
			// w.Header().Set("Content-Type", "html/text")
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "template/404Error.html")
		}
	} else {
		// w.Header().Set("Content-Type", "html/text")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	}
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if r.URL.Path != "/" && r.URL.Path != "/groupie-tracker" {
			// w.Header().Set("Content-Type", "html/text")
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "template/404Error.html")
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "template/405Error.html")
	}
	ha := r.FormValue("asciiBanner")
	
	d := struct {
		Result string
	}{
		Result: ha,
	}
	tmp, _ := template.ParseFiles("template/processor.html")

	tmp.Execute(w, d)

}

func ErrorPage(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	template.Must(template.ParseGlob("template/500Error.html"))
	sere.ExecuteTemplate(w, "template/500Error.html", nil)
}


