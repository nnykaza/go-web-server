// help.go
package main

import (
	"fmt"
	"net/http"
)

// Help handler
func helpHandler(w http.ResponseWriter, r *http.Request) {
	helpText := `
	Available APIs:
	- /hello-world      : Returns "Hello World – GWS"
	- /hello-world-html : Returns HTML "Hello World — GWS"
	- /hello-world-json : Returns JSON with message "Hello World – GWS"
	- /syllabus         : GET for syllabus info, POST for create (stubbed), PUT for update (stubbed), DELETE for remove (stubbed)
	- /help             : Returns this help message
	`
	fmt.Fprintf(w, helpText)
}
