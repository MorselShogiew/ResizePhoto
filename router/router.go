package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MorselShogiew/ResizePhoto/middleware"
)

type API interface {
	Router(r *mux.Router)
}

func Router(apis ...API) http.Handler {
	root := mux.NewRouter().StrictSlash(true).PathPrefix("/").Subrouter()

	api := root.PathPrefix("/api").Subrouter()

	// setup services' routes
	for i := range apis {
		// for auth management
		v := api.PathPrefix("/").Subrouter()
		apis[i].Router(v)
	}

	return middleware.CorsMiddleware(root)
}
