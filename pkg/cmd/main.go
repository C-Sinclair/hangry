package main

import (
	"log"
	"net/http"
	"os"

	web "github.com/c-sinclair/restic/pkg/http"
)

func main() {
	port := os.Getenv("PORT")

	// setup db client from STORAGE pkg
	webService := web.New()

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}
