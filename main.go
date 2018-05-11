package main

import (
	"os"
	"log"
	"net/http"
)


const P2pPort = "6000"

func HttpPort() string {
	return os.Getenv("HTTP_PORT")
}

func main() {

	//Starting the blockchain from hardcoded genesis block
	go initBlockchain()

	//Mux Router 
	router := NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + HttpPort(), router))
}

