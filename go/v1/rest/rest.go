package rest

import (
	s "../service"
	"fmt"
	"github.com/fukata/golang-stats-api-handler"
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
	r.HandleFunc("/stats", getStats()).
		Methods("POST")
	/*
		r.HandleFunc("/token", getToken(s)).
			Methods(method_POST)
	*/
}

func getStats() handler {
	return func(w http.ResponseWriter, req *http.Request) {
		// read input
		authorization := req.Header.Get("authorization")
		if authorization != "bearer <special token>" {
			w.WriteHeader(401)
			fmt.Fprintf(w, "Not authorized")
			return
		}
		stats_api.Handler(w, req)
	}
}
