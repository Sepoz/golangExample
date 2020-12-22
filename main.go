package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	s := newServer(":8080", mux)

	mux.HandleFunc("/", handleHi)

	s.ListenAndServe()
}

func newServer(port string, mux *http.ServeMux) *http.Server {
	s := &http.Server{
		Addr: port,
		Handler: mux,
	}
	return s
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}