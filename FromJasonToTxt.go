package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"
)

type Data struct {
	// Define struct fields based on the JSON structure
	ID   int    `json:"id"`
	Name string `json:"name"`
	Band string `json:"band"`
	// ...
}

func FromJasonToTxt(w http.ResponseWriter) {

	// Make an HTTP GET request to fetch the JSON data
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		fmt.Println("Error fetching JSON data:", err)
		ErrorPage(w)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		ErrorPage(w)
		return
	}

	// Initialize a slice to store the JSON records
	var jsonData []Data

	// Unmarshal the JSON data into the slice
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		ErrorPage(w)
		return
	}

	// Convert JSON data to bytes
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		ErrorPage(w)
		return
	}

	// Write JSON data to a file
	jsonFile, err := os.Create("../data/data.txt")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		ErrorPage(w)
		return
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonBytes)
	if err != nil {
		fmt.Println("Error writing JSON data:", err)
		ErrorPage(w)
		return
	}

	fmt.Println("JSON data fetched and saved successfully!")

}
