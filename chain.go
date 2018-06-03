package main

/* package main

import (
	"encoding/json"
	"log"
	"time"
)

//NewBlockchain append the genesis block
func NewBlockchain() *Blockchain {
	mutex.Lock()
	defer mutex.Unlock()
	return &Blockchain{0, "", []*Block{genesisBlock()}}
}

func toJSON(bc *Blockchain, latestBlockFlag ...bool) string {
	var blks = bc.blocks
	//if only latestBlock asked to be encoded
	if latestBlockFlag != nil && latestBlockFlag[0] {
		blks = []*Block{bc.blocks[len(bc.blocks)-1]}
	}
	encoded, err := json.Marshal(blks)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	return string(encoded)
}

func fromJSON(encoded string) *Blockchain {
	var bc []*Block
	err := json.Unmarshal([]byte(encoded), &bc)
	if err != nil {
		log.Fatal("Cannot decode to JSON ", err)
	}
	return &Blockchain{0, "", bc}
}

func replaceBlockchain(currentBlockchain *Blockchain, newBlockchain *Blockchain) *Blockchain {
	if isValidChain(newBlockchain) {
		return newBlockchain
	}
	return currentBlockchain
}

func addBlock(bc *Blockchain, nextBlock *Block) []*Block {
	if isValidBlock(nextBlock, latestBlock(bc)) {
		bc.blocks = append(bc.blocks, nextBlock)
	}
	return bc.blocks
}

func generateNextBlock(bc *Blockchain, data Transaction) *Block {
	nextTimestamp := uint64(time.Now().Unix())
	previousBlock := latestBlock(bc)
	nextIndex := previousBlock.Index + 1
	nextHash := calculateHash(nextIndex, previousBlock.Hash, data, nextTimestamp)
	return &Block{nextIndex, nextHash, previousBlock.Hash, data, nextTimestamp}
}

func latestBlock(bc *Blockchain) *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func isValidBlock(newBlock *Block, previousBlock *Block) bool {
	if (previousBlock.Index + 1) != newBlock.Index {
		return false
	} else if previousBlock.Hash != newBlock.PreviousHash {
		return false
	} else if calculateHashForBlock(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func isValidChain(newBlockchain *Blockchain) bool {
	if newBlockchain.blocks[0].Hash != genesisBlock().Hash {
		return false
	}

	//tempBlockchain := Blockchain{[]*Block{genesisBlock()}}
	tempBlockchain := []*Block{genesisBlock()}
	for i := 1; i < len(newBlockchain.blocks); i++ {
		if isValidBlock(newBlockchain.blocks[i], tempBlockchain[i-1]) {
			tempBlockchain = append(tempBlockchain, newBlockchain.blocks[i])
		} else {
			return false
		}
	}

	return true
}
*/
