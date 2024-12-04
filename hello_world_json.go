// hello_world_json.go
package main

import (
	"encoding/json"
	"net/http"
)

// Hello World JSON handler
func helloWorldJSONHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "Hello World – GWS",
	}
	json.NewEncoder(w).Encode(response)
}
