package main

//NodeIDType  string
type NodeIDType string

//AddressType string
type AddressType string

//LeaderType struct
type LeaderType struct {
	nextHeight  map[string]uint64
	matchHeight map[string]uint64
}

//Node States
const (
	follower  int = 0
	candidate int = 1
	leader    int = 2
	shutdown  int = 3
)

// Node represent the current operating daemon
type Node struct {
	ID          *NodeIDType
	State       int
	Leader      *LeaderType
	Address     *AddressType
	Blockchain  *Blockchain
	BlockHeight uint64
}

//NewNode creates a new node in the network initialized with hardcoded genesisBlock
func NewNode(id *NodeIDType, address *AddressType) *Node {
	return &Node{id, follower, &LeaderType{}, address, &Blockchain{0, nil, []*Block{genesisBlock()}}, 0}
}

//AppendBlock is invoked by leader to replicate blocks; also used as heartbeat
func (n *Node) AppendBlock(leaderTerm int, leaderID NodeIDType, blockHeight uint64, blockTerm int, blocks []*Block, leaderBlockHeight uint64) (int, bool) {
	//TODO
	if leaderTerm < n.Blockchain.currentTerm {
		return n.Blockchain.currentTerm, false
	}

	println("Sending request for mandatory replication")
	println(n.Blockchain.currentTerm)
	println(n.ID)
	println(n.Blockchain.currentTerm - 1)
	println(n.BlockHeight)
	println([]*Block{})
	println("nextBlock Height")

	return 0, true

}

//RequestVote is invoked by candidates to gather votes
func (n *Node) RequestVote() {
	println("Sending request for gather vote")
}
