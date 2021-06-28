package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"wordcounter/wordcounter"
)

// main initialises the word store and begins a net/http server on port 9001
func main() {
	// initialise the word store
	wordcounter.NewWordStore()

	// begin a net/http server on port 9001
	http.HandleFunc("/counts", counts)
	http.HandleFunc("/text", text)
	http.HandleFunc("/", unsupported)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}

// counts handles GET /counts requests
func counts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		wordcounter.GetCounts(w)
		return
	}
	error(w, fmt.Sprint("Only GET requests accepted at endpoint: ", r.URL.Path))
}

// text handles POST /text requests
func text(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		wordcounter.ProcessInput(w, r)
		return
	}
	error(w, fmt.Sprint("Only POST requests accepted at endpoint: ", r.URL.Path))
}

// unsupported handles requests to unsupported endpoints
func unsupported(w http.ResponseWriter, r *http.Request) {
	error(w, fmt.Sprint("Endpoint does not exist: ", r.URL.Path))
}

// error returns an error response to the client
func error(w http.ResponseWriter, message string) {
	wordcounter.ErrorResponse(w, errors.New(message))
}
