package main

import (
	"log"
	"net/http"
)

const HttpPort = "3000"
const P2pPort = "6000"

var Blockchain []Block


// TODO Add scriptSig 
// TODO Add real UTXOs (multiple outs, ins)
//Transaction data model
type Transaction struct {
	value int
	input string
	output string
}

//Block data model
// blocksize: 1 transaction
type Block struct {
	Index int
	Hash string
	PreviousHash string
	Data Transaction
	Timestamp int
}


func main() {
	//Starting the blockchain from hardcoded genesis block
	go initBlockchain()
	//Mux Router 
	router := NewRouter()
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + HttpPort, router))
}

