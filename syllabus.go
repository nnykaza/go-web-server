package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed syllabus.json
var syllabusData []byte

// Syllabus handler (serves the embedded JSON file)
func syllabusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.Write(syllabusData)
	} else if r.Method == "DELETE" {
		fmt.Fprintf(w, "deleted – stubbed")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "create-stubbed")
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "update-stubbed")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
