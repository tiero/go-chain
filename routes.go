package main

/*
import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route is a type
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

//Routes is a type
type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/ping",
		PingHandler,
	},
	Route{
		"GET",
		"/block",
		BlocksHandler,
	},
	Route{
		"POST",
		"/block",
		NewBlockHandler,
	},
	Route{
		"GET",
		"/peer",
		PeersHandler,
	},
	Route{
		"POST",
		"/peer",
		NewPeerHandler,
	},
	Route{
		"GET",
		"/ws",
		WebSocketHandler,
	},
}

//NewRouter start new router
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
*/
