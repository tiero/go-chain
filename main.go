package main

import (
	"log"
	"net/http"
)

var blockchain *Blockchain
var node *Node

func main() {

	config := generateConfigFromFlags()

	if config.isValid() {
		//Starting the blockchain from hardcoded genesis block
		//blockchain = NewBlockchain()
		node = NewNode(config)
		//heartbeat := setTimeout(config.HeartbeatTimeout)

		//Mux Router
		router := NewRouter()
		// Bind to a port and pass our router in
		log.Fatal(http.ListenAndServe(string(config.Address), router))
	} else {
		log.Panic("Error: Provide --host, --port and --id")
	}

}
