package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var blockchain *Blockchain
var node *Node
var mutex sync.Mutex

func main() {

	//Flag Commannds
	//TODO move away
	host := flag.String("host", "", "remote host")
	port := flag.String("port", "", "Port")
	id := flag.String("id", "", "Port")
	flag.Parse()

	if *id != "" && *host != "" && *port != "" {
		serverAddress := *host + ":" + *port
		//Starting the blockchain from hardcoded genesis block
		blockchain = NewBlockchain()
		node = NewNode(*id, serverAddress, follower, latestBlock(blockchain).Index)

		//Mux Router
		router := NewRouter()
		// Bind to a port and pass our router in
		log.Fatal(http.ListenAndServe(serverAddress, router))
	} else {
		log.Fatal("Error: Provide --host, --port and --id")
	}

}
