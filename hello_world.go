// hello_world.go
package main

import (
	"fmt"
	"net/http"
)

// Hello World handler (Plain Text)
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World – GWS")
}
