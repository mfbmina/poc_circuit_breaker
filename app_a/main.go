package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success")
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/failure", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Failure")
		w.WriteHeader(http.StatusInternalServerError)
	})

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
