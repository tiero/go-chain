package main

import (
	"sync"
)

var mutex sync.Mutex

func initBlockchain() {
	// we just want to make sure only one goroutine can access a variable at a time to avoid conflicts
	mutex.Lock()
	defer mutex.Unlock()
	Blockchain = append(Blockchain, genesisBlock())
}