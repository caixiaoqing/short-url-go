package api

import (
	"net/http"
	"encoding/json"
	"github.com/learning/short-url-go/model"
	"github.com/learning/short-url-go/repo"
	"github.com/learning/short-url-go/utils"
)

/*
Test with this curl command:
curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
*/
func shortenHandler(w http.ResponseWriter, r *http.Request) {

	var u model.Url
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest) //400
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //EOF  400
		return
	}
	defer r.Body.Close()

	responseForShorten(w, longToShort(u.Url))
}

func responseForShorten(w http.ResponseWriter, url string) {
	w.Header().Set(content_type, application_json)
	w.WriteHeader(http.StatusOK)

	short := model.ShortUrl{API_URL + ORIGINAL + string(url[len(API_URL):]), url}
	if err := json.NewEncoder(w).Encode(short); err != nil {
		panic(err)
	}
}

func longToShort(url string) string {
	id := repo.RepoFindUrl(url)
	if id != -1 {
		return API_URL + idToShortURL(id)
	}

	newId := repo.RepoCreateUrl(url)
	return API_URL + idToShortURL(newId)
}

func idToShortURL(id int) string {
	shortUrl := utils.Encode(id)
	for len(shortUrl) < 6 {
		shortUrl = "0" + shortUrl
	}
	return shortUrl
}
