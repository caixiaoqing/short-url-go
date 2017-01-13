package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		indexHandler,
	},
	Route{
		"Shorten",
		"POST",
		"/shorten",
		shortenHandler,
	},
	Route{
		"Original",
		"GET",
		"/original/{short_url}",
		originalHandler,
	},
	Route{
		"IndexShort",
		"GET",
		"/{short_url}",
		indexShortHandler,
	},
}
