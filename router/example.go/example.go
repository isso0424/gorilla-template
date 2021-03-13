package example

import (
	"log"
	"net/http"
)

type Example struct {}

func(route Example) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Success!!!"))

	log.Println("Success")
}

func(route Example) Method() string {
	return "GET"
}

func(route Example) Name() string {
	return "example"
}

func(route Example) Path() string {
	return "/example"
}
