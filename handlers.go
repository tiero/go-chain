package main

import (
	"log"
	"net/http"
	"encoding/json"	
)

import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type PeerBody struct {
	Port string
}

func NewPeer(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var peerBody PeerBody
	err := decoder.Decode(&peerBody)
	log.Println(peerBody)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}

	if _, e := connectToPeer(peerBody.Port); e != nil {
		log.Println(e)
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("internal server error"))
	
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("new peer at :" + peerBody.Port))

}

/*
Peer
*/
func Peer(writer http.ResponseWriter, request *http.Request) {

	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}    
	
}


/* 
NewBlock is a test-only endpoint used for development. Do not use in production!
curl -X POST http://localhost:3000/block -H 'Content-Type: application/json' \
-d '{
      "Value":100000000,
      "Input":"@satoshi",
      "Output": "@tiero"
}' 
*/
func NewBlock(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var newTransaction Transaction
	err := decoder.Decode(&newTransaction)
	log.Println(newTransaction)

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

/* 
Blocks is used to get the whole list of blocks or a specific one 
curl -X GET http://localhost:3000/block
*/
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

/* 
Ping ... Pong!
curl -X GET http://localhost:3000/ping 
*/
func Ping(writer http.ResponseWriter, request *http.Request) {
	//Pong
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Pong!"))
}



