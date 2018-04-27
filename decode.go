package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

type httpServer struct {
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Stop here if its Preflighted OPTIONS request
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Action, Module")
	}

	if r.Method == "OPTIONS" {
		return
	}

	r.ParseForm()
	if len(r.Form["v"]) > 0 {

		w.Write([]byte(("hello") + " " + r.Form["v"][0]))
	}

}

func main() {
	addr := flag.String("http-address", "0.0.0.0:80", "")
	flag.Parse()

	var h httpServer

	httpListener, err := net.Listen("tcp", *addr)
	server := http.Server{
		Handler: &h,
	}
	server.Serve(httpListener)

	fmt.Println("finish ", err)
}
