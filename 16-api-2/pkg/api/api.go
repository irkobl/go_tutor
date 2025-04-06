package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type API struct {
	router *mux.Router
	store  *sessions.CookieStore
}

func New(r *mux.Router) *API {
	api := API{
		router: r,
		store:  sessions.NewCookieStore([]byte("secret_password")),
	}
	return &api
}

func (api *API) Endpoints() {
	api.router.Use(requestIDMiddleware)
	api.router.Use(logMiddleware)
	api.router.Use(api.jwtMiddleware)
	//api.router.Use(api.sessionsMiddleware)

	// api.router.HandleFunc("/api/v1/authSession", api.authSession).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/authJWT", api.authJWT).Methods(http.MethodPost, http.MethodOptions)
	// api.router.HandleFunc("/api/v1/books", api.books).Methods(http.MethodGet, http.MethodOptions)
	// api.router.HandleFunc("/api/v1/newBook", api.newBook).Methods(http.MethodPost, http.MethodOptions)
}
