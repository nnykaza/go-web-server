// gws.go
package main

import (
	"log"
	"net/http"
)

func main() {
	// Register handlers
	http.HandleFunc("/hello-world", helloWorldHandler)
	http.HandleFunc("/hello-world-html", helloWorldHTMLHandler)
	http.HandleFunc("/hello-world-json", helloWorldJSONHandler)
	http.HandleFunc("/syllabus", syllabusHandler)
	http.HandleFunc("/help", helpHandler)

	// Start the server
	log.Println("Server starting on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe failed: ", err)
	}
}
