package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/fetch-metar", FetchMetarHandler)
	http.ListenAndServe(":6327", nil)
}

func FetchMetarHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://metar.vatsim.net/all"

	response, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch data from URL", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, "Failed to read response body", http.StatusInternalServerError)
			return
		}

		if SaveToFile("metar.txt", body) {
			fmt.Fprintf(w, "Data fetched and saved successfully")
		} else {
			http.Error(w, "Failed to save the content to file", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Failed to fetch data from URL", http.StatusInternalServerError)
	}
}

func SaveToFile(filename string, content []byte) bool {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return false
	}
	return true
}
