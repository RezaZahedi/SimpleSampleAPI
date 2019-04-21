package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Middleware is defined solely for chaining.
type Middleware func(http.HandlerFunc) http.HandlerFunc

//Logging logs the requested URL and its process time.
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, ": Process Duration", time.Since(start))
			}()
			f(w, r)
		}
	}
}

// Splitting retrives the desired word from the request's URL path.
func (bf *BookFetcher) Splitting(totLen, index int, isInt bool) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			id, err := urlParser(r.URL.Path, totLen, index, isInt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// /////////////////////////
			// Incomplete Type Assertion, right now it works only with "isInt = true".
			// /////////////////////////
			idInt, ok := id.(int)
			if !ok {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			ctx := context.WithValue(r.Context(), ContextKeyID, idInt)
			log.Println(id, "in splitting middleware")
			f(w, r.WithContext(ctx))
		}
	}
}

// Chain is used for chaining multiple, if any, middlewares.
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// urlParser takes the URL's supposed length, the desired postion index and the type at that position.
func urlParser(sURL string, totLen, index int, isInt bool) (interface{}, error) {
	wantStr := strings.Split(sURL, "/")
	if len(wantStr) != totLen {
		return 0, fmt.Errorf("Bad URL")
	}
	if index > totLen {
		return 0, fmt.Errorf("Bad URL")
	}

	if isInt {
		wantInt, err := strconv.Atoi(wantStr[index-1])
		if err != nil {
			return 0, fmt.Errorf("Bad URL")
		}
		return wantInt, nil
	}
	return wantStr, nil
}
