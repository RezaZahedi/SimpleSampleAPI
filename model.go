package main

import (
	"sync"
)

// Book holds each element of the database structure.
type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
}

// BookFetcher is used to hold and pass database info.
type BookFetcher struct {
	session map[int]Book
	rwmutex *sync.RWMutex
}

// ContextKey is used in request's context to be compatible with the go linter suggestions.
type ContextKey string

// ContextKeyID helps to avoid key collisions in r.context.
var ContextKeyID = ContextKey("ID")

// NewBookFetcher helps creating database config.
func NewBookFetcher(m map[int]Book, rwm *sync.RWMutex) *BookFetcher {
	return &BookFetcher{m, rwm}
}
