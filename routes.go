//
// Route struct, Router implementation, Route mapping
//
package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct and Router implementation
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

//Mapping the routes to the handlers
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetAllBlogPosts",
		"GET",
		"/posts",
		GetAllBlogPosts,
	},
	Route{
		"CreateBlogPost",
		"POST",
		"/post",
		CreateBlogPost,
	},
}
