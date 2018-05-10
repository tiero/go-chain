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
// blocksize: 1 transaction per block
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

	txData := Transaction{ 
		BlockRewardValue(),
		CoinbaseInput,
		"@tiero",
	}

	return Block{ 0, GenesisBlockHash, "0", txData, GenesisTimestamp}
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

func BlockRewardValue() int { 
	return InitialBlockReward
}