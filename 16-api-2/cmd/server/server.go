package main

import (
	"gc/16-api-2/pkg/api"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = api.New(srv.router)
	srv.api.Endpoints()
	http.ListenAndServe(":8080", srv.router)
}
