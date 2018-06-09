package main

import (
	"log"
	"sync"
	"time"
)

// LeaderType struct
type LeaderType struct {
	leaderID    NodeIDType
	nextHeight  map[NodeIDType]uint64
	matchHeight map[NodeIDType]uint64
}

// Node States
const (
	follower  int = 0
	candidate int = 1
	leader    int = 2
	shutdown  int = 3
)

// Node represent the current operating daemon
type Node struct {
	Config Config
	State  int
	Leader LeaderType
	// lastContact is the last time we had contact from the
	// leader node. This can be used to gauge staleness.
	lastContact     time.Time
	lastContactLock sync.RWMutex
	//Blockchain
	Blockchain  *Blockchain
	BlockHeight uint64
}

// NewNode creates a new node in the network initialized with hardcoded genesisBlock
func NewNode(config Config) *Node {
	return &Node{config, follower, LeaderType{}, time.Time{}, sync.RWMutex{}, &Blockchain{0, "", []*Block{}}, 0}
}

func (n *Node) setState(nextState int) {
	n.State = nextState
}

func (n *Node) setLastContact() {
	n.lastContactLock.Lock()
	n.lastContact = time.Now()
	n.lastContactLock.Unlock()
}

func (n *Node) init() {
	//TODO
	//Add timeout for election and heartbeat
	heartbeatTimeout := setTimeout(n.Config.HeartbeatTimeout)
	for {
		select {
		case <-heartbeatTimeout:
			//Restart timeout
			heartbeatTimeout = setTimeout(n.Config.HeartbeatTimeout)
			if time.Now().Sub(n.lastContact) < n.Config.HeartbeatTimeout {
				continue
			}
			log.Println("The leader is lost")
			//TODO Start new election phase
			// set leaderID to zero value and become candidate
			n.Leader.leaderID = ""
			n.setState(candidate)

		}

	}
}

// AppendBlock is invoked by leader to replicate blocks; also used as heartbeat
// prevBlockHeight and prevBlockTerm defines the last known to be committed block in the leader chain store
// NextBlockHeight is the leader last commited block Index
func (n *Node) AppendBlock(leaderTerm int, leaderID NodeIDType, prevBlockHeight uint64, prevBlockTerm int, nextBlockHeight uint64, nextBlock *Block) (int, bool) {
	//TODO Create helper for rules
	//If request or response contains term T > currentTerm: set currentTerm = T, convert to follower
	if leaderTerm > n.Blockchain.currentTerm {
		n.Blockchain.currentTerm = leaderTerm
		n.setState(follower)
	}

	//save the leaderID so we can redirects the clients
	n.Leader.leaderID = leaderID
	// Reply false if term < currentTerm
	if leaderTerm < n.Blockchain.currentTerm {
		return n.Blockchain.currentTerm, false
	}

	//Reply false if log doesn’t contain an entry at prevLogIndex whose term matches prevLogTerm
	if n.Blockchain.getBlockAtIndex(prevBlockHeight) == nil {
		return n.Blockchain.currentTerm, false
	}
	if n.Blockchain.getBlockAtIndex(prevBlockHeight).Term != prevBlockTerm {
		return n.Blockchain.currentTerm, false
	}
	//If an existing block conflicts with a new one (same index but different terms), delete the existing block and all that follow it
	if n.Blockchain.getBlockAtIndex(nextBlockHeight) != nil && n.Blockchain.getBlockAtIndex(nextBlockHeight).Term != nextBlock.Term {
		n.Blockchain.removeBlocksFromIndex(nextBlockHeight)
	}

	//Append any block after checking if is valid against the previous one
	if nextBlock != nil {
		if isValidBlock(nextBlock, n.Blockchain.getBlockAtIndex(n.BlockHeight)) {
			n.Blockchain.addBlock(nextBlock)
			//Advance the index of highest block known to be committed
			n.BlockHeight = nextBlockHeight
			return n.Blockchain.currentTerm, true
		}
		return n.Blockchain.currentTerm, false
	}

	//It is only an heartbeat, do all the election timeout stuff
	log.Println("heartbeat: we should restart the election timeout")
	//prevBlockHeight and prevLogTerm match, regardless if it is
	//has been added or not a new block we return tru
	n.setLastContact()
	return n.Blockchain.currentTerm, true

}

//RequestVote is invoked by candidates to gather votes
func (n *Node) RequestVote(candidateTerm int, candidateID NodeIDType, blockHeight uint64, blockTerm int) (int, bool) {
	println("Sending request for gather vote")

	if candidateTerm < n.Blockchain.currentTerm {
		return n.Blockchain.currentTerm, false
	}
	if n.BlockHeight > blockHeight || n.Blockchain.currentTerm > blockTerm {
		return n.Blockchain.currentTerm, false
	}
	//Grant a vote if null or already equal to candidateID
	if n.Blockchain.votedFor == "" || n.Blockchain.votedFor == candidateID {
		return n.Blockchain.currentTerm, true
	}

	// This means Already voted for someone else in the current term
	// no possibility to grant a vote for the election
	return n.Blockchain.currentTerm, false
}
