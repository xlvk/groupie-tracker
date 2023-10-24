package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type GroupieData struct {
	Artists  string `json:"artists"`
	Location string `json:"location"`
	Dates    string `json:"dates"`
	Relation string `json:"relation"`
}

func main() {

}

func groupieHandler(w http.ResponseWriter, r *http.Request) {
	// Locate and Read JSON File
	fileData := "GroupieData.json"
	data, err := ioutil.ReadFile(fileData)
	if err != nil {
		fmt.Println("Error: ", err)
	}

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
