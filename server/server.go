package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Server is running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
