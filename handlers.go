package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type peerBody struct {
	Peers []string
}

/*
Peer Get a list of connected peers
curl -X GET http://localhost:3000/peer
*/
func Peer(writer http.ResponseWriter, request *http.Request) {
	response, err := json.MarshalIndent(node.Peers, "", "  ")
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
NewPeer Add an array of peers with the JSON format { "Peers":["ws://localhost:4000"] }
*/
func NewPeer(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var body peerBody
	err := decoder.Decode(&body)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}

	connected := connectToPeers(node, body.Peers)

	log.Println(connected)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("new peers!"))

}

/*
WebSocketHandler upgrade http to ws for incoming connection
*/
func WebSocketHandler(writer http.ResponseWriter, request *http.Request) {
	upgrader.CheckOrigin = func(request *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		println(err.Error())
		return
	}
	wsListen(node, conn)
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

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad Request"))
		return
	}

	//Add the next block using goroutine
	nextBlock := generateNextBlock(blockchain, newTransaction)
	go addBlock(blockchain, nextBlock)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("new block mined! Block Hash: " + nextBlock.Hash))
}

/*
Blocks is used to get the whole list of blocks or a specific one
curl -X GET http://localhost:3000/block
*/
func Blocks(writer http.ResponseWriter, request *http.Request) {
	response, err := json.MarshalIndent(blockchain.blocks, "", "  ")
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
