package lilGroupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetData(w http.ResponseWriter, r *http.Request) ([]artistsAPI, error) {
	generalData, err := http.Get("https://groupietrackers.herokuapp.com/api/")

	if err != nil {
		fmt.Print(err.Error())
		fmt.Println(9)
		ErrorPage(w, r)
		return nil, err
	}
	// fmt.Println(generalData)
	defer generalData.Body.Close()
	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		// fmt.Print(err.Error())
		// ErrorPage(w, r)
		fmt.Println(10)

		return nil, err
		// return
	}
	defer fileData.Body.Close()

	// fileData := ""
	data2, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		// fmt.Println("Error1: ", err)
		// ErrorPage(w, r)
		fmt.Println(11)

		return nil, err
		// return
	}
	var groupieData []artistsAPI
	err = json.Unmarshal(data2, &groupieData)
	if err != nil {
		fmt.Println("Error12: ", err)
		// ErrorPage(w, r)
		fmt.Println(12)

		return nil, err
		// return
	}
	return groupieData, nil
}

func GetRelations(id int, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")

	if err != nil {
		fmt.Print(err.Error())
		fmt.Println(13)
		ErrorPage(w, r)
	}
	defer fileData.Body.Close()
	data, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		fmt.Println("Error1: ", err)
		fmt.Println(14)
		ErrorPage(w, r)
	}

	var groupieData index1
	err = json.Unmarshal(data, &groupieData)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Print(len(groupieData.Index))
	
	
	return groupieData.Index[0].Relations, nil
}

func GetLocations(id int, w http.ResponseWriter, r *http.Request) ([]string, error) {
	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		fmt.Print(err.Error())
		fmt.Println(16)
		ErrorPage(w, r)
	}
	defer fileData.Body.Close()
	data, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		fmt.Println("Error3: ", err)
		fmt.Println(17)

		ErrorPage(w, r)
	}
	var groupieData []LocationsAPI
	err = json.Unmarshal(data, &groupieData)
	if err != nil {
		fmt.Println("Error4: ", err)
		fmt.Println(18)

	}
	return groupieData[id-1].Locations, nil
}

func GetDates(id int, w http.ResponseWriter, r *http.Request) ([]string, error) {
	fileData, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if err != nil {
		fmt.Print(err.Error())
		fmt.Println(19)

		ErrorPage(w, r)
	}
	defer fileData.Body.Close()
	data, err := ioutil.ReadAll(fileData.Body)
	if err != nil {
		fmt.Println("Error5: ", err)
		fmt.Println(20)
		ErrorPage(w, r)
	}
	var groupieData []DateAPI
	err = json.Unmarshal(data, &groupieData)
	if err != nil {
		fmt.Println("Error6: ", err)
		fmt.Println(21)
	}
	return groupieData[id-1].Dates, nil
}
