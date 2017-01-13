package main

import (
	"github.com/caixiaoqing/short-url-go/api"
	"github.com/caixiaoqing/short-url-go/repo"
	"log"
	"net/http"
)

func main() {
	repo.InitRepo()
	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
