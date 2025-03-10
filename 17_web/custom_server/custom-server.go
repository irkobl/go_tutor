package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	const addr = ":8080"
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Handler:      nil,
		Addr:         addr,
	}
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", mainHandler)

	log.Fatal(srv.Serve(listener))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h2>Go Simple Web App</h2></body></html>`))
}
