package main

import (
	"net/http"
	"encoding/json"	
)

func Blocks(writer http.ResponseWriter, request *http.Request) {
	response, err := json.MarshalIndent(Blockchain, "", "  ")
	//Catch the error
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(response))
}

func Ping(writer http.ResponseWriter, request *http.Request) {
	//Pong
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Pong"))
}



