package main

import (
	"time"
)

// TODO Add scriptSig
// TODO Add real UTXOs (multiple outs, ins)

//Transaction data model
type Transaction struct {
	Value  uint64
	Input  string
	Output string
}

//Block data model
// blocksize: 1 transaction per block
type Block struct {
	Index        uint64
	Hash         string
	PreviousHash string
	Data         Transaction
	Timestamp    uint64
}

//Blockchain data model
type Blockchain struct {
	blocks []*Block
}

func genesisBlock() *Block {

	txData := Transaction{
		BlockRewardValue(),
		CoinbaseInput,
		"@tiero",
	}

	return &Block{0, GenesisBlockHash, "0", txData, GenesisTimestamp}
}

//NewBlockchain append the genesis block
func NewBlockchain() *Blockchain {
	mutex.Lock()
	defer mutex.Unlock()
	return &Blockchain{[]*Block{genesisBlock()}}
}

func replaceBlockchain(currentBlockchain *Blockchain, newBlockchain *Blockchain) *Blockchain {
	if isValidChain(newBlockchain) {
		return newBlockchain
	}
	return currentBlockchain
}

func addBlock(bc *Blockchain, nextBlock *Block) {
	if isValidBlock(nextBlock, latestBlock(bc)) {
		mutex.Lock()
		defer mutex.Unlock()
		bc.blocks = append(bc.blocks, nextBlock)
	}
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

//BlockRewardValue returns the next BlockReward
func BlockRewardValue() uint64 {
	return InitialBlockReward
}
