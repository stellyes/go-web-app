package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")

		if err != nil {
			fmt.Println(fmt.Sprintf("%s", err))
		}

		fmt.Println(fmt.Sprintf("Number of bytes written %d", n))
	})

	// Start web server (port, and handler)
	_ = http.ListenAndServe(":4000", nil)
}
