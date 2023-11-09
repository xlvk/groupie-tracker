package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	// "os"
	"strconv"
	// "strings"
)

var sere *template.Template

func init() {
	sere = template.Must(template.ParseGlob("index.html"))
}

// type GroupieData struct {
// 	artists   string `json:"artists"`
// 	Locations string `json:"locations"`
// 	Dates     string `json:"dates"`
// 	Relation  string `json:"relation"`
// }

var groupieData []artistsAPI

// var generalData []GroupieData

type GroupieData struct {
	artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type artistsAPI struct {
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
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", groupieHandler)
	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	// http.HandleFunc("/groupie-tracker", processor)
	http.HandleFunc("/aboutme.html", aboutMe)
	// submit?value=2
	http.HandleFunc("/artist", artistPage)
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

func groupieHandler(w http.ResponseWriter, r *http.Request) {
	// Locate and Read JSON File
	if r.URL.Path != "/" && r.URL.Path != "/aboutme.html" && r.URL.Path != "/artist" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")
	}
	groupieData, err := GetData(w, r)
	if err != nil {
		// Handle the error
		ErrorPage(w, r)
		return
	}

	// Prepare Data For HTML
	type DataView struct {
		Groupie []artistsAPI
	}
	viewData := DataView{
		Groupie: groupieData,
	}

	// Load HTML Page
	pageView := template.Must(template.ParseFiles("index.html"))

	// Execute HTML with Data
	err = pageView.Execute(w, viewData)
	if err != nil {
		fmt.Println("Error3: ", err)
		ErrorPage(w, r)
		return
	}

}
func aboutMe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/aboutme.html" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")

	}
	http.ServeFile(w, r, "aboutme.html")
}

func artistPage(w http.ResponseWriter, r *http.Request) {
	groupieData, err := GetData(w, r)
	if err != nil {
		// Handle the error
		ErrorPage(w, r)
		return
	}

	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil || id <= 0 || id > len(groupieData) {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")
		// http.Redirect(w, r, "/400", http.StatusSeeOther)
		return
	}

	var err2 error
	fmt.Print("\n" + groupieData[id-1].Name + "\n")
	groupieData[id-1].Relations, err2 = GetRelations(id, w, r)
	if err2 != nil {
		// Handle the error
		ErrorPage(w, r)
		return
	}

	temp, err2 := template.ParseFiles("artistpage.html")
	if err2 != nil {
		// Handle the error
		ErrorPage(w, r)
		return
	}

	err3 := temp.Execute(w, groupieData[id-1])
	if err3 != nil {
		// Handle the error
		ErrorPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func GetRelations(id int, w http.ResponseWriter, r *http.Request) (string, error) {
	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		fmt.Print(err.Error())
		// os.Exit(1)
		ErrorPage(w, r)
		// return
	}
	defer fileData.Body.Close()

	// fileData := ""
	data, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		fmt.Println("Error1: ", err)
		ErrorPage(w, r)
		// return
	}
	var groupieData []artistsAPI
	err = json.Unmarshal(data, &groupieData)
	if err != nil {
		fmt.Println("Error2: ", err)
		ErrorPage(w, r)
		// return  err
		// return
	}
	// fmt.Print(groupieData[0].Members)""
	// fmt.Print("\n" + groupieData[id].Name + "\n")
	// fmt.Print(groupieData[0].Locations)
	return groupieData[id-1].Relations, nil
}

func GetData(w http.ResponseWriter, r *http.Request) ([]artistsAPI, error) {
	generalData, err := http.Get("https://groupietrackers.herokuapp.com/api/")

	if err != nil {
		fmt.Print(err.Error())
		ErrorPage(w, r)
		return nil, err
	}
	defer generalData.Body.Close()

	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		// fmt.Print(err.Error())
		// ErrorPage(w, r)
		return nil, err
		// return
	}
	defer fileData.Body.Close()

	// fileData := ""
	data2, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		// fmt.Println("Error1: ", err)
		// ErrorPage(w, r)
		return nil, err
		// return
	}
	var groupieData []artistsAPI
	err = json.Unmarshal(data2, &groupieData)
	if err != nil {
		// fmt.Println("Error2: ", err)
		// ErrorPage(w, r)
		return nil, err
		// return
	}
	return groupieData, nil
}
