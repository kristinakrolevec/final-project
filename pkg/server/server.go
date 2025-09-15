package server

import (
	"FINAL-PROJECT/pkg/api"
	"log"
	"net/http"
)

const webDir = "./web"

var MyServer *http.Server

func Run() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	api.Init(mux)

	MyServer := &http.Server{
		Addr:    ":7540",
		Handler: mux,
	}

	log.Printf("Start server")

	err := MyServer.ListenAndServe()
	if err != nil {
		log.Fatal("Mistake in ListenAndServe")
	}

}
