package main

import (
	"fmt"
	"net/http"
	"vet/pets/internal"
)

func main() {

	http.HandleFunc("/", internal.Pets)
	http.HandleFunc("/{id}", internal.Pet)
	fmt.Println("[ pets microservice ] listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
