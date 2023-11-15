package lilGroupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AboutMe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/aboutme.html" {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "template/404Error.html")

	}
	http.ServeFile(w, r, "template/aboutme.html")
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	groupieData, err := GetData(w, r)
	if err != nil {
		// Handle the error
		fmt.Println(22)
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
		ErrorPage(w, r)
		return
	}

	temp, err2 := template.ParseFiles("template/artistpage.html")
	if err2 != nil {
		// Handle the error
		fmt.Println(24)
		ErrorPage(w, r)
		return
	}

	err3 := temp.Execute(w, groupieData[id-1])
	if err3 != nil {
		// Handle the error
		fmt.Println(25)
		ErrorPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
