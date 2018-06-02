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
