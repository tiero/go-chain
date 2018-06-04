package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Handlers
func pingHandler(writer http.ResponseWriter, request *http.Request) {
	//Pong
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Pong!"))
}

func blocksHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := json.MarshalIndent(node.Blockchain.blocks, "", "  ")
	//Catch the error
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(response))
}

func addPeersHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var body peerBody
	err := decoder.Decode(&body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}

	//Add peers
	//TODO

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("new peers!"))
}

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
		pingHandler,
	},
	Route{
		"GET",
		"/block",
		blocksHandler,
	},
	Route{
		"POST",
		"/peer",
		addPeersHandler,
	},

	/*
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
		}, */
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
