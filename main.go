package main

import (
	"log"
	"net/http"
)

const HttpPort = "3000"
const P2pPort = "6000"


func main() {
	//Starting the blockchain from hardcoded genesis block
	go initBlockchain()
	//Mux Router 
	router := NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + HttpPort, router))
}

