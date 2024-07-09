package internal

import (
	"encoding/json"
	"net/http"
	"vet/pets/pkg"
	"vet/pets/repository"
)

func Pet(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getOne(w, r)
	case http.MethodPut:
		update(w, r)
	case http.MethodDelete:
		delete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getOne(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pkg.ErrorMessage{StatusCode: http.StatusBadRequest, Message: ErrorMissingId})
		return
	}

	doc, present := repository.GetOne(id)
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add(HeaderContentType, ContentTypeJSON)
	json.NewEncoder(w).Encode(doc)
}

func update(w http.ResponseWriter, r *http.Request) {

	var doc pkg.Pet
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pkg.ErrorMessage{StatusCode: http.StatusBadRequest, Message: ErrorMissingId})
		return
	}
	doc.Id = id
	repository.Update(doc)

	w.Header().Add(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(doc)
}

func delete(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pkg.ErrorMessage{StatusCode: http.StatusBadRequest, Message: ErrorMissingId})
		return
	}

	err := repository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
