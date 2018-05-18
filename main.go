package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var blockchain *Blockchain
var node *Node

var mutex sync.Mutex

//HTTPPort Exposes curent http port
func HTTPPort() string {
	return os.Getenv("HTTP_PORT")
}

//Host Exposes current Host
func Host() string {
	return "localhost:" + HTTPPort()
}
func main() {

	//Starting the blockchain from hardcoded genesis block
	blockchain = NewBlockchain()
	node = NewNode(Host(), []*websocket.Conn{}, latestBlock(blockchain).Index)
	//Mux Router
	router := NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(Host(), router))
}
