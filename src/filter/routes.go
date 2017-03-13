package filter

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"FilterIndex",
		"GET",
		"/filters",
		FilterIndex,
	},
	Route{
		"FilterShow",
		"GET",
		"/filters/{filterId}",
		FilterShow,
	},
	Route{
		"FilterCreate",
		"POST",
		"/filters",
		FilterCreate,
	},
}
