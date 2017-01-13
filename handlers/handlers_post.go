package handlers

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
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) //EOF
		return
	}
	defer r.Body.Close()

	responseForShorten(w, longToShort(u.Url))
}

func responseForShorten(w http.ResponseWriter, url string) {
	w.Header().Set(content_type, application_json)
	w.WriteHeader(http.StatusOK)

	short := model.ShortUrl{api_url + original + string(url[len(api_url):]), url}
	if err := json.NewEncoder(w).Encode(short); err != nil {
		panic(err)
	}
}

func longToShort(url string) string {
	id := repo.RepoFindUrl(url)
	if id != -1 {
		return api_url + idToShortURL(id)
	}

	newId := repo.RepoCreateUrl(url)
	return api_url + idToShortURL(newId)
}

func idToShortURL(id int) string {
	shortUrl := utils.Encode(id)
	for len(shortUrl) < 6 {
		shortUrl = "0" + shortUrl
	}
	return shortUrl
}
