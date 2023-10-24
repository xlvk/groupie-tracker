package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var sere *template.Template

func init() {
	sere = template.Must(template.ParseGlob("index.html"))
}

// type GroupieData struct {
// 	Artists   string `json:"artists"`
// 	Locations string `json:"locations"`
// 	Dates     string `json:"dates"`
// 	Relation  string `json:"relation"`
// }

type GroupieData struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type ArtistsAPI struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsAPI struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

func main() {
	http.HandleFunc("/", groupieHandler)
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/groupie-tracker", processor)
	u, err := url.Parse("http://localhost:2003")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Listening and serving on: ")
	fmt.Printf("%+v", u)
	fmt.Println()
	log.Fatal(http.ListenAndServe(":2003", nil))
}

func groupieHandler(w http.ResponseWriter, r *http.Request) {
	// Locate and Read JSON File

	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer fileData.Body.Close()

	// fileData := ""
	data, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	//fmt.Println(string(data))
	// Parse JSON Data
	var groupieData []GroupieData
	err = json.Unmarshal(data, &groupieData)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Prepare Data For HTML
	type DataView struct {
		Groupie []GroupieData
	}
	viewData := DataView{
		Groupie: groupieData,
	}

	// Load HTML Page
	pageView := template.Must(template.ParseFiles("index.html"))

	// Execute HTML with Data
	err = pageView.Execute(w, viewData)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Print Data on Terminal
	// for _, group := range groupieData {
	// 	fmt.Println(group.Artists, group.Dates, group.Location, group.Relation)
	// }
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		if r.URL.Path != "/" && r.URL.Path != "/groupie-tracker" {
			// w.Header().Set("Content-Type", "html/text")
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "404Error.html")
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "405Error.html")
	}
	ha := r.FormValue("asciiBanner")

	d := struct {
		Result string
	}{
		Result: ha,
	}
	tmp, _ := template.ParseFiles("processor.html")

	tmp.Execute(w, d)

}
