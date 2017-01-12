package main

import (
	"log"
	"net/http"
	"github.com/learning/short-url-go/handlers"
)

func main() {
	//initRepo()
	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
