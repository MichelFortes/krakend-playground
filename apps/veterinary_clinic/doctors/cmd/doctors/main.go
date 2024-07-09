package main

import (
	"fmt"
	"net/http"
	"vet/doctors/internal"
)

func main() {

	http.HandleFunc("/", internal.Doctors)
	http.HandleFunc("/{id}", internal.Doctor)
	fmt.Println("[ doctors microservice ] listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
