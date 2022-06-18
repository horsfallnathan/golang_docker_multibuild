package main

import (
	"log"
	"net/http"
	"time"
)

type healthHandler struct {
}

func (hh healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server Running"))
}

func main() {
	m := http.NewServeMux()
	m.Handle("/healthz", healthHandler{})

	s := &http.Server{
		Handler: m,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port 8080")
	log.Fatal(s.ListenAndServe())
}
