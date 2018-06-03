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
	votedFor    *NodeIDType
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

//BlockRewardValue returns the next BlockReward
func BlockRewardValue() uint64 {
	return InitialBlockReward
}
