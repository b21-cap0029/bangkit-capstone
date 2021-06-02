package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/handler"
	"github.com/b21-cap0029/bangkit-capstone/cc/aaida-backend/internal/models"
)

const (
	defaultBindAddress = "0.0.0.0:8080"
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
	server.HandleFunc("/health", handler.Health)
	server.Handle("/check", handler.NewDefaultCheckHandler())

	err := models.ConnectDataBase()
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe(bindAddress, server)
	if err != nil {
		log.Fatalln(err)
	}
}
