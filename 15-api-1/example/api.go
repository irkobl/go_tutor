package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func (api *API) Endpoints() {
	api.router.Use(logMiddleware)
	api.router.Use(headersMiddleware)
	api.router.HandleFunc("/api/v1/books", api.books).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/books", api.newBook).Methods(http.MethodPost, http.MethodGet)
}

func (api *API) books(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) newBook(w http.ResponseWriter, r *http.Request) {
	var b book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	books = append(books, b)
}
