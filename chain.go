package main

import (
	"sync"
	"time"
)


// TODO Add scriptSig 
// TODO Add real UTXOs (multiple outs, ins)
//Transaction data model
type Transaction struct {
	Value int
	Input string
	Output string
}

//Block data model
// blocksize: 1 transaction
type Block struct {
	Index int
	Hash string
	PreviousHash string
	Data Transaction
	Timestamp int64
}

// Our beloved and complicated blockchain <3
var Blockchain []Block

var mutex sync.Mutex

func genesisBlock() Block {
	// Wednesday 9th May 2018 10:16:19 PM UTC
	return Block{0, "3cd45a480c2601ed55245eac8b233c680f111eaad30c568a318e5213f7f0f522", "0", Transaction{}, 1525904179}
}

func latestBlock() Block {
	return Blockchain[len(Blockchain)-1]
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
	Blockchain = newBlockchain
}

func addBlock(nextBlock Block) {
	mutex.Lock()
	defer mutex.Unlock()
	if isValidBlock(nextBlock, latestBlock()) {
		Blockchain = append(Blockchain, nextBlock)
	}
}

func generateNextBlock(data Transaction) Block {
	nextTimestamp 	:= time.Now().Unix()
	previousBlock 	:= latestBlock()
	nextIndex 	:= previousBlock.Index + 1
	nextHash 	:= calculateHash(nextIndex, previousBlock.Hash, data, nextTimestamp)
	return Block{ nextIndex, nextHash, previousBlock.Hash, data, nextTimestamp }
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