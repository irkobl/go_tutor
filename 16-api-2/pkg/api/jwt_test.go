package api

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func TestAPI_authJWT(t *testing.T) {

	type fields struct {
		router *mux.Router
		store  *sessions.CookieStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				router: tt.fields.router,
				store:  tt.fields.store,
			}
			api.authJWT(tt.args.w, tt.args.r)
		})
	}
}
