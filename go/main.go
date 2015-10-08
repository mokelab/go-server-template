package main

import (
	"./v1"
	"flag"
	"fmt"
	"github.com/fukata/golang-stats-api-handler"
	"github.com/gorilla/mux"
	"github.com/mokelab-go/hupwriter"
	"log"
	"net"
	"net/http"
)

const (
	AppName = "myapp"
)

func main() {
	r := mux.NewRouter()
	err := v1.InitRouter(r, log.New(hupwriter.New("/var/log/"+AppName+".log", "/var/pid/"+AppName+".pid"), "logger:", log.Lshortfile))
	if err != nil {
		fmt.Printf("Failed to init server : %s", err)
		return
	}
	var standalone *bool = flag.Bool("standalone", false, "if true, runs standalone mode")
	flag.Parse()

	if *standalone {
		initStandalone(r)
	}

	initStatAPI(r)

	l, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Printf("Failed to call net.Listen : %s", err)
		return
	}

	fmt.Println("server is running on 9001")
	//err = fcgi.Serve(l, r)
	err = http.Serve(l, r)
	if err != nil {
		fmt.Printf("failed to stop : %s", err)
	}
}

func initStandalone(r *mux.Router) {
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("../web/"))))
}

func initStatAPI(r *mux.Router) {
	r.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")
		if authorization != "bearer bLuNUD44a5TUcMVBFczTQnyDQxk3ALiuSxAN2ZwxzRUYUc5m" {
			w.WriteHeader(401)
			fmt.Fprintf(w, "Not authorized")
			return
		}
		stats_api.Handler(w, r)
	}).Methods("GET")
}
