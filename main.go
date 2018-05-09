package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/hex"
	"encoding/json"
	"crypto/sha256"

	"github.com/gorilla/mux"
)


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

//Router
func main() {
	//Start genesis block
	Blockchain = append(Blockchain, genesisBlock())

	//Mux Router 
	router := mux.NewRouter()
	// Http Endpoint 
	router.HandleFunc("/", Ping)
	router.HandleFunc("/blocks", Blocks)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Endpoints
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

// helpers
func calculateHash(index int, prevHash string, data Transaction, timestamp int64) string {
	payload := fmt.Sprint(index) + prevHash + fmt.Sprint(data.value) + data.input + data.output + fmt.Sprint(timestamp)
	h := sha256.New()
	h.Write([]byte(payload))

	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)	
}

func genesisBlock() Block {
	// Wednesday 9th May 2018 10:16:19 PM UTC
	return Block{0, "3cd45a480c2601ed55245eac8b233c680f111eaad30c568a318e5213f7f0f522", "0", Transaction{}, 1525904179}
}

/*
func Hash(writer http.ResponseWriter, request *http.Request) {
	t := time.Now().Unix()
	//OK
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(calculateHash(0, "0", Transaction{}, t)))
	fmt.Println(fmt.Sprint(t))
}
*/