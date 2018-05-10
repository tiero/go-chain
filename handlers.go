package main

import (
	"net/http"
	"encoding/json"	
)



func NewBlock(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var newTransaction Transaction
	err := decoder.Decode(&newTransaction)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
		return	
	}

	//Add the next block using goroutine
	nextBlock := generateNextBlock(newTransaction)
	go addBlock(nextBlock)


	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("new block mined! Block Hash: " + nextBlock.Hash))
}

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



