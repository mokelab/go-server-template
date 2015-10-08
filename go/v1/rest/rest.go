package rest

import (
	s "../service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	method_GET    = "GET"
	method_POST   = "POST"
	method_PUT    = "PUT"
	method_DELETE = "DELETE"
)

type handler func(http.ResponseWriter, *http.Request)

func SetHandlers(r *mux.Router, services s.Services) {
	/*
		r.HandleFunc("/token", getToken(s)).
			Methods(method_POST)
	*/
}
