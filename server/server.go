package server

import (
	"isso0424/gorilla-template/router"
	"net/http"

	"github.com/gorilla/mux"
)

func serve(c *Config) error {
	r := mux.NewRouter()

	for _, route := range router.Routes {
		handler := createHandler(route)

		r.Methods(route.Method()).Path(route.Path()).Name(route.Name()).Handler(handler)
	}

	return nil
}

func createHandler(route router.Route) http.Handler {
	handler := route

	// TODO: add midleware in here
	return handler
}
