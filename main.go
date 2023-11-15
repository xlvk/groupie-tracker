package main

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"lilGroupie/lilgroupie"

	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "os"
	// "strconv"
	// "strings"
)

var Sere *template.Template

func init() {
	Sere = template.Must(template.ParseGlob("template/index.html"))
}

func main() {
	// http.Handle("/template/css/", http.StripPrefix("/template/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/", lilGroupie.GroupieHandler)
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	// http.HandleFunc("/groupie-tracker", processor)
	http.HandleFunc("/aboutme.html", lilGroupie.AboutMe)
	// submit?value=2
	http.HandleFunc("/artist", lilGroupie.ArtistPage)
	// http.Handle("/pics/", http.FileServer(http.Dir("templates")))

	// http.Handle("/css/", http.FileServer(http.Dir("templates")))
	u, err := url.Parse("http://localhost:2003")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Listening and serving on: ")
	fmt.Printf("%+v", u)
	fmt.Println()
	log.Fatal(http.ListenAndServe(":2003", nil))
}
