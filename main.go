package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	//http://localhost:8080
	router.HandleFunc("/", indexHandler)

	//http://localhost:8080/shorten
	router.HandleFunc("/shorten", shortenHandler)

	//http://localhost:8080/original/{short_url}
	router.HandleFunc("/original/{shorten_url}", originalHandler)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	//data := parsePostBody
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func originalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! original")
}
