package server

import (
	"fmt"
	"isso0424/gorilla-template/router"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve(c *Config) error {
	r := mux.NewRouter()

	for _, route := range router.Routes {
		handler := createHandler(route)

		r.Methods(route.Method()).Path(route.Path()).Name(route.Name()).Handler(handler)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", c.ListenPort), handlers.CompressHandler(r))
}

func createHandler(route router.Route) (handler http.Handler) {
	handler = handlers.LoggingHandler(os.Stdout, http.HandlerFunc(route.ServeHTTP))

	return
}
