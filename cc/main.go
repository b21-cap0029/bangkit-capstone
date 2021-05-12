package main

import (
	"flag"
	"log"
	"net/http"

	handler "github.com/b21-cap0029/bangkit-capstone/tree/master/cc/internal/handler"
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
	server.HandleFunc("/", handler.Index)
	server.HandleFunc("/health", handler.Health)
	http.ListenAndServe(bindAddress, server)
}
