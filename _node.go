/* package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//MessageType struct
const (
	queryLatestBlock   int = 0
	queryAllBlock      int = 1
	queryState         int = 2
	responseBlockchain int = 3
	responseAck        int = 4
	appendNextBlock    int = 5
	heartbeatPing      int = 6
)

//MessagePayload struct
type MessagePayload struct {
	SenderHost  string
	SenderState int
	MessageType int
	MessageText string
}

// Node represent the current operating daemon
type Node struct {
	ID      string
	Address string
	Term    int
	State   int
	Peers   map[string]*websocket.Conn
	Leader  string
	Height  uint64
}

func (n *Node) logState() string {
	switch n.State {
	case follower:
		return "Follower"
	case candidate:
		return "Candidate"
	case leader:
		return "Leader"
	case shutdown:
		return "Shutdown"
	default:
		return "Unknown"
	}
}

// NewNode starts a p2p server
func NewNode(id, host string, state int, blocksHeight uint64) *Node {
	mutex.Lock()
	defer mutex.Unlock()
	return &Node{id, host, 0, state, map[string]*websocket.Conn{}, "", blocksHeight}
}

func appendBlockProposal(n *Node, nextBlock *Block) bool {

	//NOTICE Only for dev purpose, remove it
	//TODO Create an election mechanism to handle leader-follower writes
	if n.State != leader {
		n.setState(leader)
	}
	//Actual check
	if n.isLeader() {
		blocks := addBlock(blockchain, nextBlock)
		n.Height = nextBlock.Index
		broadcastMessage(node, &MessagePayload{n.Address, n.State, responseBlockchain, toJSON(&Blockchain{blockchain.currentTerm, blockchain.votedFor, blocks})})
	}

	return n.isLeader()
}

func (n *Node) isLeader() bool {
	return n.State == leader
}

func (n *Node) setState(nextState int) {
	n.State = nextState
}

func connectToPeers(n *Node, endpoints []string) {

	for _, endpoint := range endpoints {
		//Parse endpoints
		// TODO move to helper function
		parsedURL, _ := url.Parse(endpoint)
		u := url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host, Path: "/ws"}
		//Start a dialing
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

		if err == nil {
			go wsListen(n, c)
			broadcastMessage(n, &MessagePayload{n.Address, n.State, queryLatestBlock, ""})
		} else {
			log.Println(err)
		}
	}
}

func wsListen(n *Node, conn *websocket.Conn) {
	n.Peers[n.Address] = conn
	println(conn.RemoteAddr().String() + " Connected")
	println("Peers: " + fmt.Sprint(len(n.Peers)))
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			println(conn.RemoteAddr().String() + " Disconnected")
			removePeer(n, conn)
			println("Peers: " + fmt.Sprint(len(n.Peers)))
			return
		}
		handleIncomingMessage(n, p)
	}
}

func removePeer(n *Node, conn *websocket.Conn) bool {
	_, ok := n.Peers[n.Address]
	if ok {
		delete(n.Peers, n.Address)
		return true
	}
	return false
}

func broadcastMessage(n *Node, mp *MessagePayload) {
	payload, err := json.Marshal(mp)
	for _, connection := range n.Peers {
		if err == nil {
			if err = connection.WriteMessage(websocket.TextMessage, payload); err != nil {
				return
			}
		}
	}
}

func handleResponseBlockchain(n *Node, message string) {
	receivedBlockchain := fromJSON(message)
	receivedBlockchainLenght := len(receivedBlockchain.blocks)
	latestBlockReceived := receivedBlockchain.blocks[receivedBlockchainLenght-1]
	latestBlockHeld := latestBlock(blockchain)

	if latestBlockReceived.Index > latestBlockHeld.Index {
		log.Println("blockchain possibly behind. We got: " + fmt.Sprint(latestBlockHeld.Index) + " Peer got: " + fmt.Sprint(latestBlockReceived.Index))
		if latestBlockHeld.Hash == latestBlockReceived.PreviousHash {
			log.Println("We can safely append the chain")
			//broadcastMessage(n, &MessagePayload{node.State, responseAck, ""})
			addBlock(blockchain, latestBlockReceived)
			n.Height = latestBlockReceived.Index
			broadcastMessage(n, &MessagePayload{n.Address, n.State, responseBlockchain, toJSON(blockchain, true)})
		} else if receivedBlockchainLenght == 1 {
			//In this case we should check if we are at genesis block
			log.Println("We have to query the chain from our peer")
			broadcastMessage(n, &MessagePayload{n.Address, n.State, queryAllBlock, ""})
		} else {
			log.Println("Received blockchain is longer than current blockchain")
			blockchain = replaceBlockchain(blockchain, receivedBlockchain)
			n.Height = latestBlock(receivedBlockchain).Index
		}
	} else {
		log.Println("Received blockchain is not longer than current blockchain. Do nothing")
	}

}

func handleResponseAck(n *Node, message string) {
	latestBlockHeld := latestBlock(blockchain)
	//Do things
	println(latestBlockHeld)
}

func handleQueryState(n *Node, mp MessagePayload) {
	//totalNumPeers := len(n.Peers)

}

func handleHeartbeatPing(n *Node, mp MessagePayload) {
	//filterConnectionFromEndpoint(n.Peers, mp.SenderHost)
	if mp.SenderState == leader {
		//do nothing
	}
}

func handleIncomingMessage(n *Node, message []byte) {
	var payload MessagePayload
	if err := json.Unmarshal(message, &payload); err == nil {
		switch int(payload.MessageType) {
		case queryState:
			handleQueryState(n, payload)
			//broadcastMessage(n, &MessagePayload{n.Address, n.State, queryLatestBlock, ""})
		case queryLatestBlock:
			broadcastMessage(n, &MessagePayload{n.Address, n.State, responseBlockchain, toJSON(blockchain, true)})
		case queryAllBlock:
			broadcastMessage(n, &MessagePayload{n.Address, n.State, responseBlockchain, toJSON(blockchain)})
		case responseBlockchain:
			handleResponseBlockchain(n, payload.MessageText)
		case responseAck:
			handleResponseAck(n, payload.MessageText)
		case heartbeatPing:
			handleHeartbeatPing(n, payload)
		}

	} else {
		log.Println(err)
	}
}
*/