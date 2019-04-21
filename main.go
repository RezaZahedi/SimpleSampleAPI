package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	rwmutex := sync.RWMutex{}
	bf := NewBookFetcher(getSession(), &rwmutex)

	// show all database
	http.Handle("/books", Chain(bf.ReadAll, Logging()))

	// handles query parameters
	http.Handle("/book", Chain(bf.Query, Logging()))

	//handles GET, POST and DELETE
	http.Handle("/book/", Chain(bf.MethodMux, bf.Splitting(3, 3, true), Logging()))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// getSession initializes the map data structure, it is supposed to be like a database.
func getSession() map[int]Book {
	return map[int]Book{
		1: {1, "Revelation", "des R"},
		2: {2, "Elevation", "des E"},
		3: {3, "Zensation", "des Z"},
		4: {4, "Annihilation", "des A"},
		6: {6, "Elevation", "des E2"},
	}
}
