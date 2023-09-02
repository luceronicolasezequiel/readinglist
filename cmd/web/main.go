package main

import (
	"flag"
	"log"
	"net/http"

	"readinglist.luceronicolasezequiel/internal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:4000/v1/books", "Endpoint for the reading list web service")

	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
