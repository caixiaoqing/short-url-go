package api_test

import (
	"net/http/httptest"
	"io"
	"fmt"
	"encoding/json"
	"testing"
	"strings"
	"net/http"

	"github.com/caixiaoqing/short-url-go/api"
	"github.com/caixiaoqing/short-url-go/model"
	"github.com/caixiaoqing/short-url-go/repo"
)

const test_long_url_1 = "http://a.very.long.url.1"
const test_long_url_2 = "http://a.very.long.url.2"
const test_redirect_url_good = ""
const test_redirect_url_fail = "a.very.long.url.redirect"

func assertEqual(t *testing.T, expected interface{}, actual interface{}, message string) {
	if expected == actual {
		return
	}

	assertFailedMsg := fmt.Sprintf("Expected: %v != %v", expected, actual)
	message = fmt.Sprintf("%q : %s", assertFailedMsg, message)
	t.Fatal(message)
}

var (
	server   	*httptest.Server
	reader   	io.Reader
	shortenUrl 	string
	originalUrl	string
	redirect	string
)

func init() {
	repo.InitRepo()
	server = httptest.NewServer(api.NewRouter()) //Creating new server with the handlers
	shortenUrl = fmt.Sprintf("%s/shorten", server.URL) //Grab the address for the API endpoint
	originalUrl = fmt.Sprintf("%s/original/", server.URL)
	redirect = fmt.Sprintf("%s/", server.URL)
}

//TestCase: Create Shorten Url
func TestShortenHandler(t *testing.T) {
	//Step 1 : send request for "POST /shorten"
	res, err := createShortenUrl(test_long_url_1)
	if err != nil {
		t.Errorf("Error while sending request POST /shorten %s", err)
		return
	}
	//Step 2 : verify the response
	_ = parseResponsePOSTShorten(t, res)
}

//TestCase: Get Original Url from short url
func TestOriginalHandler(t *testing.T) {
	//Step 1 : create shorten url by "POST /shorten"
	resOfPost, err := createShortenUrl(test_long_url_2)
	if err != nil {
		t.Errorf("Error while sending request POST /shorten %s", err)
		return
	}

	//Step 2 : parse the response from "POST /shorten"
	u := parseResponsePOSTShorten(t, resOfPost)

	//Step 3 : send request for "GET /original/{short}"
	original := originalUrl + u.Short[len(api.API_URL + api.ORIGINAL):]
	request, err := http.NewRequest("GET", original, nil)
	res, err := http.DefaultClient.Do(request)

	//Step 4 : verify the response
	assertEqual(t, http.StatusOK, res.StatusCode, "StatusCode of response is not as expected")

	var longOri model.OriginalUrl
	err = json.NewDecoder(res.Body).Decode(&longOri)
	if err != nil {
		t.Errorf("Error while decode the response body from GET /original/{short} %s", err)
	}

	assertEqual(t, test_long_url_2, longOri.Original, "This long url from 'GET /original/{short}' doesn't match from the original")
}

//TestCase: Redirect to Original Url from short url
//	    To make it succeed, let the original url be the address of the root of the test server
func TestIndexShortHandler(t *testing.T) {
	//Step 1 : create shorten url by "POST /shorten"
	resOfPost, err := createShortenUrl(redirect + test_redirect_url_good)
	if err != nil {
		t.Errorf("Error while sending request POST /shorten %s", err)
		return
	}

	//Step 2 : parse the response from "POST /shorten"
	u := parseResponsePOSTShorten(t, resOfPost)

	//Step 3 : send request for "GET /{short}"
	redirectUrl := redirect + u.Short[len(api.API_URL + api.ORIGINAL):]

	request, err := http.NewRequest("GET", redirectUrl, nil)
	res, err := http.DefaultClient.Do(request)

	//Step 3 : verify the response (200)
	if err != nil {
		t.Errorf("Error while sending request Get /{short} %s", err)
	}

	assertEqual(t, http.StatusOK, res.StatusCode, "StatusCode of response is not as expected")
}

//TestCase: Redirect to Original Url from short url
//	    Off course it will fail if the original url is not existing
func TestIndexShortHandlerFailed(t *testing.T) {
	//Step 1 : create shorten url by "POST /shorten"
	resOfPost, err := createShortenUrl(redirect + test_redirect_url_fail)
	if err != nil {
		t.Errorf("Error while sending request POST /shorten %s", err)
		return
	}

	//Step 2 : parse the response from "POST /shorten"
	u := parseResponsePOSTShorten(t, resOfPost)

	//Step 3 : send request for "GET /{short}"
	redirectUrl := redirect + u.Short[len(api.API_URL + api.ORIGINAL):]

	request, err := http.NewRequest("GET", redirectUrl, nil)
	res, err := http.DefaultClient.Do(request)

	//Step 3 : verify the response (400)
	if err != nil {
		t.Errorf("Error while sending request Get /{short} %s", err)
	}
	assertEqual(t, http.StatusBadRequest, res.StatusCode, "StatusCode of response is not as expected")
}

func createShortenUrl(url string) (*http.Response, error) {
	u := &model.Url{url}
	jsonStr, err := json.Marshal(u)

	reader = strings.NewReader(string(jsonStr)) //Convert string to reader

	request, err := http.NewRequest("POST", shortenUrl, reader) //Create request with JSON body
	res, err := http.DefaultClient.Do(request)
	return res, err
}

func parseResponsePOSTShorten(t *testing.T, r *http.Response) model.ShortUrl {

	assertEqual(t, http.StatusOK, r.StatusCode, "StatusCode of response is not as expected")

	var u model.ShortUrl
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		t.Errorf("Error while decode the response body from POST /shorten %s", err)
	}
	return u;
}
