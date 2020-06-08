package main

import (
	"log"
	"net/http"
	"os"

	web "github.com/c-sinclair/restic/pkg/http"
	"github.com/c-sinclair/restic/pkg/storage"
)

func main() {
	port := os.Getenv("PORT")

	store := storage.New()
	webService := web.New(store)

	log.Printf("Running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, webService.Router))
}
