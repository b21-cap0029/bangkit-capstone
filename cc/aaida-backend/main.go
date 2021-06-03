package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

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
	err := models.ConnectDataBase()
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health)
	router.Handle("/check", handler.NewDefaultCheckHandler())

	subCases := router.PathPrefix("/cases/").Subrouter()
	subCases.Handle("/submit", handler.NewDefaultCasesSubmitHandler())
	subCases.Handle("/{id:[0-9]+}", handler.NewDefaultCasesHandler())
	subCases.Handle("/{id:[0-9]+}/claim", handler.NewDefaultCasesClaimHandler())
	subCases.Handle("/{id:[0-9]+}/skip", handler.NewDefaultCasesSkipHandler())
	subCases.Handle("/{id:[0-9]+}/close", handler.NewDefaultCasesCloseHandler())

	subProfile := router.PathPrefix("/profile").Subrouter()
	subProfile.Handle("/", handler.NewDefaultProfileHandler())

	n := negroni.Classic()
	n.UseHandler(router)

	err = http.ListenAndServe(bindAddress, n)
	if err != nil {
		log.Fatalln(err)
	}
}
