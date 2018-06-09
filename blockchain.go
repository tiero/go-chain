package main

//Transaction data model
// TODO Add scriptSig
// TODO Add real UTXOs (multiple outs, ins)
type Transaction struct {
	Value  uint64
	Input  string
	Output string
}

//Block data model | blocksize: 1 transaction per block
type Block struct {
	Index        uint64
	Term         int
	Hash         string
	PreviousHash string
	Data         Transaction
	Timestamp    uint64
}

//Blockchain data model
type Blockchain struct {
	currentTerm int
	votedFor    NodeIDType
	blocks      []*Block
}

func genesisBlock() *Block {

	txData := Transaction{
		BlockRewardValue(),
		CoinbaseInput,
		"@tiero",
	}

	return &Block{0, 1, GenesisBlockHash, "0", txData, GenesisTimestamp}
}

func (bc *Blockchain) getBlockAtIndex(index uint64) *Block {
	for _, b := range bc.blocks {
		if b.Index == index {
			return b
		}
	}
	return nil
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

func (bc *Blockchain) addBlock(nextBlock *Block) {
	bc.blocks = append(bc.blocks, nextBlock)
}

func (bc *Blockchain) removeBlocksFromIndex(index uint64) {
	for i, b := range bc.blocks {
		if b.Index == index {
			bc.blocks = bc.blocks[:i]
		}
	}
}

//BlockRewardValue returns the next BlockReward
func BlockRewardValue() uint64 {
	return InitialBlockReward
}
