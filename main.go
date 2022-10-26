package main

import (
	"fmt"
	"net/http"
)

var version = "0.0.6"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Test 5, WebServer is listening on 8080")
	})

	http.ListenAndServe(":8080", nil)
}
