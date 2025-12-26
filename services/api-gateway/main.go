package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API Gateway...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from API Gateway!"))
	})

	http.ListenAndServe(":8080", nil)
}
