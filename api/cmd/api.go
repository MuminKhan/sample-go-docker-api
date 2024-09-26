package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server at port 8001")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Fprintf(w, "Hello, World!")
	})

	if err := http.ListenAndServe(":8001", nil); err != nil {
		fmt.Println(err)
	}
}
