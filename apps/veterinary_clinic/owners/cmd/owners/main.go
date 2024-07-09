package main

import (
	"fmt"
	"net/http"
	"vet/owners/internal"
)

func main() {

	http.HandleFunc("/", internal.Owners)
	http.HandleFunc("/{id}", internal.Owner)
	fmt.Println("[ owners microservice ] listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
