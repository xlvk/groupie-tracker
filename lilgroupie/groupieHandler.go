package lilGroupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func GroupieHandler(w http.ResponseWriter, r *http.Request) {
	// Locate and Read JSON File
	if r.URL.Path != "/" && r.URL.Path != "/aboutme.html" && r.URL.Path != "/artists/" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")
	}
	groupieData, err := GetData(w, r)
	if err != nil {
		// Handle the error
		fmt.Println(1)
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
	pageView := template.Must(template.ParseFiles("template/index.html"))

	// Execute HTML with Data
	err = pageView.Execute(w, viewData)
	if err != nil {
		fmt.Println("Error3: ", err)
		fmt.Println(2)
		ErrorPage(w, r)
		return
	}

}
