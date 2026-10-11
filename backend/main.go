package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the server on port 8080
	println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
