package handlers

import (
	"net/http"
	"fmt"
	"html"
	"encoding/json"
	"github.com/learning/short-url-go/model"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

/*
Test with this curl command:
curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
*/
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	//data := parsePostBody
	var u model.Url
	if r.Body == nil {
		http.Error(w, "Please send a request body",  400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)  //EOF
		return
	}
	defer r.Body.Close()
	fmt.Fprintf(w, "parsed : %q\n", u.Url)
}

func originalHandler(w http.ResponseWriter, r *http.Request) {
	//data := parsePostBody
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
