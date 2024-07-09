package internal

import (
	"encoding/json"
	"net/http"
	"vet/pets/pkg"
	"vet/pets/repository"

	"github.com/google/uuid"
)

func Pets(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		create(w, r)
	case http.MethodGet:
		getAll(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func create(w http.ResponseWriter, r *http.Request) {

	var doc pkg.Pet
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pkg.ErrorMessage{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	doc.Id = uuid.New().String()
	repository.Create(doc)

	w.Header().Add(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(doc)
}

func getAll(w http.ResponseWriter, _ *http.Request) {

	docs := repository.GetAll()

	w.Header().Add(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(docs)
}
