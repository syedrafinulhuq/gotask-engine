package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GoTask Server Started on :8080")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":9090", nil)
}
