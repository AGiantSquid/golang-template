package handler

import (
	"fmt"
	"net/http"
)

// HelloHandler returns a simple hello message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Call the GreetHandler with the name
	fmt.Fprint(w, "Hello, World!")
}
