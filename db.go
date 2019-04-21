package main

import "encoding/json"

// BookError is used in reporting database related user errors.
type BookError struct {
	S string
}

func (e *BookError) Error() string {
	return string(e.S)
}

func ReadBooks(bf *BookFetcher) ([]byte, error) {
	bf.rwmutex.RLock()
	defer bf.rwmutex.RUnlock()

	return json.Marshal(bf.session)
}

func ReadBook(bf *BookFetcher, id int) ([]byte, error) {
	bf.rwmutex.RLock()
	defer bf.rwmutex.RUnlock()

	book, ok := bf.session[id]
	if !ok {
		return nil, &BookError{"Book Not Found"}
	}
	return json.Marshal(book)
}

func CreateBook(bf *BookFetcher, id int, book Book) error {
	bf.rwmutex.Lock()
	defer bf.rwmutex.Unlock()

	_, ok := bf.session[id]
	if ok {
		return &BookError{"Bad Request: Book Already Exists"}
	}
	if book.ID != id {
		return &BookError{"Bad Request: Missmatching IDs"}
	}

	bf.session[book.ID] = book
	return nil
}

func DeleteBook(bf *BookFetcher, id int) error {
	bf.rwmutex.Lock()
	defer bf.rwmutex.Unlock()

	_, ok := bf.session[id]
	if !ok {
		return &BookError{"Bad Request: Book Not Found"}
	}
	delete(bf.session, id)
	return nil
}

func QueryBook(bf *BookFetcher, name string) ([]byte, error) {
	var books []Book
	bf.rwmutex.RLock()
	defer bf.rwmutex.RUnlock()

	for _, book := range bf.session {
		if book.Name == name {
			books = append(books, book)
		}
	}
	if books == nil {
		return nil, &BookError{"Found No Book Named: " + name}
	}
	return json.Marshal(books)
}
