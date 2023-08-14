package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Display a list of the book on the reading list")
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Added a new book to the reading list")
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBook(w, r)
	case http.MethodPut:
		app.updateBook(w, r)
	case http.MethodDelete:
		app.deleteBook(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (app *application) getBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Display the details of book with ID: %d", idInt)
}

func (app *application) updateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Update the details of book with ID: %d", idInt)
}

func (app *application) deleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Delete a book with ID: %d", idInt)
}
