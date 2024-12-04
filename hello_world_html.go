// hello_world_html.go
package main

import (
	"fmt"
	"net/http"
)

// Hello World HTML handler
func helloWorldHTMLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World – GWS</h1>")
}
