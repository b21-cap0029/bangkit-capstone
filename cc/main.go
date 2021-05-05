package main

import (
	"log"
	"flag"
	"net/http"
)

const (
	defaultBindAddress = "0.0.0.0:8000"
)

func main() {
	bindAddress := flag.String("bind", defaultBindAddress, "Address to listen")
	flag.Parse()

	log.Print("========================= AAIDA =========================")
	log.Printf("Server started at %s\n", *bindAddress)

	serveHTTP(*bindAddress)
}

func serveHTTP(bindAddress string) {
	server := http.NewServeMux()
	server.HandleFunc("/", index)
	server.HandleFunc("/status", status)
	http.ListenAndServe(bindAddress, server)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Hello, World"}`))
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Healthy"}`))
}
