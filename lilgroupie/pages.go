package lilGroupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AboutMe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/aboutme.html" {
		ErrorPage(w, r, 404)
		return

	}
	http.ServeFile(w, r, "template/aboutme.html")
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	groupieData, err := GetData(w, r)
	if err != nil {
		// Handle the error
		fmt.Println(22)
		ErrorPage(w, r, 500)
		return
	}

	sid := r.URL.Query().Get("id")
	id, err := strconv.Atoi(sid)
	if err != nil || id <= 0 || id > len(groupieData) {
		ErrorPage(w, r, 404)
		// http.Redirect(w, r, "/400", http.StatusSeeOther)
		return
	}

	var err2 error
	fmt.Print("\n" + groupieData[id-1].Name + "\n")
	// arr, err2 :=
	// ha := ""
	// var row = 1

	// var result = []string{}
	// for _, column := range arr {
	// 	result = append(result, column[row])
	// }
	// for i := 0; i < len(result); i++ {
	// 	ha += result[i] + "\n"
	// }
	groupieData[id-1].Relations, err2 = GetRelations(id, w, r)
	if err2 != nil {
		// Handle the error
		fmt.Println(23)
		ErrorPage(w, r, 500)
		return
	}

	temp, err2 := template.ParseFiles("template/artistpage.html")
	if err2 != nil {
		// Handle the error
		fmt.Println(24)
		ErrorPage(w, r, 500)
		return
	}

	err3 := temp.Execute(w, groupieData[id-1])
	if err3 != nil {
		// Handle the error
		fmt.Println(25)
		ErrorPage(w, r, 500)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
