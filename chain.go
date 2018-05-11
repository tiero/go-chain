package main

import (
	"net/url"
	"log"
	"sync"
	"time"
)

import "github.com/gorilla/websocket"

// TODO Add scriptSig 
// TODO Add real UTXOs (multiple outs, ins)

//Transaction data model
type Transaction struct {
	Value uint64
	Input string
	Output string
}

//Block data model
// blocksize: 1 transaction per block
type Block struct {
	Index uint64
	Hash string
	PreviousHash string
	Data Transaction
	Timestamp uint64
}

// Our beloved and complicated blockchain <3
var Blockchain []Block
var Peers []string

var mutex sync.Mutex

func genesisBlock() Block {

	txData := Transaction{ 
		BlockRewardValue(),
		CoinbaseInput,
		"@tiero",
	}

	return Block{ 0, GenesisBlockHash, "0", txData, GenesisTimestamp}
}


func connectToPeer(port string) (*websocket.Conn, error)  {
	u := url.URL{Scheme: "ws", Host:"localhost:" + port, Path: "/peer"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	
	if err == nil {
		go addPeer(u.String())
	}

	return c, err
}

func addPeer(endpoint string) {
	mutex.Lock()
	defer mutex.Unlock()
	Peers = append(Peers, endpoint)
}


func initBlockchain() {
	// we just want to make sure only one goroutine can access a variable at a time to avoid conflicts
	mutex.Lock()
	defer mutex.Unlock()
	Blockchain = append(Blockchain, genesisBlock())
}



func replaceBlockchain(newBlockchain []Block) {
	mutex.Lock()
	defer mutex.Unlock()
	if isValidChain(newBlockchain) {
		Blockchain = newBlockchain
	}
}

func addBlock(nextBlock Block) {
	mutex.Lock()
	defer mutex.Unlock()
	if isValidBlock(nextBlock, latestBlock()) {
		Blockchain = append(Blockchain, nextBlock)
	}
}

func generateNextBlock(data Transaction) Block {
	nextTimestamp 	:= uint64(time.Now().Unix())
	previousBlock 	:= latestBlock()
	nextIndex 	:= previousBlock.Index + 1
	nextHash 	:= calculateHash(nextIndex, previousBlock.Hash, data, nextTimestamp)
	return Block{ nextIndex, nextHash, previousBlock.Hash, data, nextTimestamp }
}

func latestBlock() Block {
	return Blockchain[len(Blockchain)-1]
}

func isValidBlock(newBlock Block, previousBlock Block) bool {
	if ((previousBlock.Index + 1) != newBlock.Index) {
		return false
	} else if previousBlock.Hash != newBlock.PreviousHash {
		return false
	} else if (calculateHashForBlock(newBlock) != newBlock.Hash) {
		return false 
	}
	return true
}

func isValidChain(newBlockchain []Block) bool {
	if newBlockchain[0].Hash != genesisBlock().Hash {
		return false
	}

	tempBlockchain := []Block{newBlockchain[0]}
	for i := 1; i < len(newBlockchain); i++ {
		if isValidBlock(newBlockchain[i], tempBlockchain[i-1]) {
			tempBlockchain = append(tempBlockchain, newBlockchain[i]) 
		} else {
			return false
		}
	}

	return true
}

func BlockRewardValue() uint64 { 
	return InitialBlockReward
}