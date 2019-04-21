package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ReadAll writes all the data as json.
func (bf *BookFetcher) ReadAll(w http.ResponseWriter, r *http.Request) {
	data, err := ReadBooks(bf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(data))

	// curl localhost:8080/books
}

// MethodMux determines the requests method to decide further actions.
func (bf *BookFetcher) MethodMux(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bf.ReadOne(w, r)
		return
	case "POST":
		bf.CreateOne(w, r)
		return
	case "DELETE":
		bf.DeleteOne(w, r)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// ReadOne returns the book data as json if it finds it.
func (bf *BookFetcher) ReadOne(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(ContextKeyID).(int) // error already checked in middlewares.go
	// log.Println(id)
	data, err := ReadBook(bf, id)
	if _, ok := err.(*BookError); ok {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(data))

	// curl localhost:8080/book/2
}

// CreateOne checks the book's ID, if it doesnt exist it adds it to the database and redirects to the "/books" URL.
func (bf *BookFetcher) CreateOne(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(ContextKeyID).(int) // error already checked in middlewares.go

	book := Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = CreateBook(bf, id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)

	//	curl -X POST -d '{"Id":5, "Name":"Reza"}' -L localhost:8080/book/5
}

// DeleteOne checks the book's ID, if it does exist it deletes it from the database and redirects to the "/books" URL.
func (bf *BookFetcher) DeleteOne(w http.ResponseWriter, r *http.Request) {
	id, _ := r.Context().Value(ContextKeyID).(int) // error already checked in middlewares.go

	err := DeleteBook(bf, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)

	//	curl -X DELETE -L localhost:8080/book/5
}

// Query searches the database and returns all the matching instances based on only name attribute.
func (bf *BookFetcher) Query(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"]
	if len(name) != 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	data, err := QueryBook(bf, name[0])
	if _, ok := err.(*BookError); ok {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(data))

	//	curl localhost:8080/book?name="Elevation"
}
