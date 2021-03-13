package midleware

import (
	"fmt"
	"isso0424/gorilla-template/logger"
	"isso0424/gorilla-template/router"
	"net/http"
)

func RegisterLogger(route router.Route) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		logger.InfoLogger.Println(
			fmt.Sprintf("%s %s %s", route.Method(), route.Path(), route.Name()),
		)

		route.ServeHTTP(w, r)
	})
}