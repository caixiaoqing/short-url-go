package main

import (
	"github.com/learning/short-url-go/handlers"
	"github.com/learning/short-url-go/repo"
	"log"
	"net/http"
)

func main() {
	repo.InitRepo()
	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
