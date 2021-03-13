package router

import "net/http"

type Route interface {
	Name() string
	Path() string
	Method() string
	Handlefunc(writer http.ResponseWriter, request http.Request)
}
