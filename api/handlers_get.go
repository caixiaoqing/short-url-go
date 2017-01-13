package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/learning/short-url-go/repo"
	"github.com/learning/short-url-go/utils"
	"github.com/learning/short-url-go/status"
	"github.com/learning/short-url-go/model"
)

// Test with this curl command:
// curl -sX GET -H 'Content-Type: application/json' 'localhost:8080'
//
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Index")
	//TODO usage template
}

// Test with this curl command:
// curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original/000001'
//
func originalHandler(w http.ResponseWriter, r *http.Request) {
	//Validate Short URL
	shortUrl := r.URL.Path[len("/original/"):]
	if !utils.IsValid(shortUrl) {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Short Url is invalid.", http.StatusBadRequest) //400
		return
	}
	//Look-up original long url from short
	responseForOriginal(w, repo.RepoFindUrlById(utils.Decode(shortUrl)))
}

// Test with this curl command:
// curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/000001'
//
func indexShortHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]

	//Validate Short URL
	if !utils.IsValid(shortUrl) {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Short Url is invalid.", http.StatusBadRequest) //400
		return
	}

	//Look-up original long url from short and redirect to it
	redirect(w, r, repo.RepoFindUrlById(utils.Decode(shortUrl)))
}

func responseForOriginal(w http.ResponseWriter, url string) {
	w.Header().Set(content_type, application_json)

	if len(url) == 0 {
		w.WriteHeader(http.StatusNotFound)
		urlShortenErr := status.HTTPError{http.StatusNotFound, url_not_found}
		if err := json.NewEncoder(w).Encode(urlShortenErr); err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	original := model.OriginalUrl{url}
	if err := json.NewEncoder(w).Encode(original); err != nil {
		panic(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set(content_type, application_json)

	if len(url) == 0 {
		w.WriteHeader(http.StatusNotFound)
		urlShortenErr := status.HTTPError{http.StatusNotFound, url_not_found}
		if err := json.NewEncoder(w).Encode(urlShortenErr); err != nil {
			panic(err)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

