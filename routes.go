package main

import (
	"net/http"
	"github.com/gorilla/mux"	
)

type Route struct {
	Method	string
	Path	string
	Handler http.HandlerFunc
} 

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/ping",
		Ping,
	},
	Route{
		"GET",
		"/block",
		Blocks,
	},
	Route{ 
		"POST", 
		"/block",
		NewBlock,
	},
	{
		"GET",
		"/peer",
		Peer,
	},
	{
		"POST",
		"/peer",
		NewPeer,
	},
}


func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Handler(route.Handler)
	}

	return router
	
}