package server

import (
	"FINAL-PROJECT/pkg/api"
	"log"
	"net/http"
	"os"
)

const webDir = "./web"

func Run() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(webDir)))

	api.Init(mux)

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	Server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Start server")

	err := Server.ListenAndServe()
	if err != nil {
		log.Fatal("Mistake in ListenAndServe")
	}

}
