package main

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
	responseBlockchain int = 2
)

//States
const (
	follower  int = 0
	candidate int = 1
	leader    int = 2
	shutdown  int = 3
)

//MessagePayload struct
type MessagePayload struct {
	MessageType int
	MessageText string
}

// Node represent the current operating daemon
type Node struct {
	Host   string
	State  int
	Peers  []*websocket.Conn
	Height uint64
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
func NewNode(host string, state int, initialPeers []*websocket.Conn, blocksHeight uint64) *Node {
	mutex.Lock()
	defer mutex.Unlock()
	return &Node{host, state, initialPeers, blocksHeight}
}

func appendBlockProposal(n *Node, nextBlock *Block) bool {

	if isLeader(n) {
		blocks := addBlock(blockchain, nextBlock)
		broadcastMessage(node, &MessagePayload{responseBlockchain, toJSON(&Blockchain{blocks})})
	}

	return isLeader(n)
}

func isLeader(n *Node) bool {
	return n.State == leader
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
			broadcastMessage(n, &MessagePayload{queryLatestBlock, ""})
		} else {
			log.Println(err)
		}
	}
}

func wsListen(n *Node, conn *websocket.Conn) {
	n.Peers = append(n.Peers, conn)
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

func removePeer(node *Node, conn *websocket.Conn) bool {
	for i, v := range node.Peers {
		if v == conn {
			node.Peers = append(node.Peers[:i], node.Peers[i+1:]...)
			return true
		}
	}
	return false
}

func broadcastMessage(n *Node, mp *MessagePayload) {
	for _, connection := range n.Peers {
		payload, err := json.Marshal(mp)
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
	println(n.Host)
	latestBlockReceived := receivedBlockchain.blocks[receivedBlockchainLenght-1]
	latestBlockHeld := latestBlock(blockchain)

	if latestBlockReceived.Index > latestBlockHeld.Index {
		log.Println("blockchain possibly behind. We got: " + fmt.Sprint(latestBlockHeld.Index) + " Peer got: " + fmt.Sprint(latestBlockReceived.Index))
		if latestBlockHeld.Hash == latestBlockReceived.PreviousHash {
			log.Println("We can safely append the new block to our chain")
			addBlock(blockchain, latestBlockReceived)
			broadcastMessage(n, &MessagePayload{responseBlockchain, toJSON(blockchain, true)})
		} else if receivedBlockchainLenght == 1 {
			//In this case we should check if we are at genesis block
			log.Println("We have to query the chain from our peer")
			broadcastMessage(n, &MessagePayload{queryAllBlock, ""})
		} else {
			log.Println("Received blockchain is longer than current blockchain")
			replaceBlockchain(blockchain, receivedBlockchain)
		}
	} else {
		log.Println("Received blockchain is not longer than received blockchain. Do nothing")
	}

}

func handleIncomingMessage(n *Node, message []byte) {
	var payload MessagePayload
	if err := json.Unmarshal(message, &payload); err == nil {
		switch int(payload.MessageType) {
		case queryLatestBlock:
			//latest block
			broadcastMessage(n, &MessagePayload{responseBlockchain, toJSON(blockchain, true)})
		case queryAllBlock:
			broadcastMessage(n, &MessagePayload{responseBlockchain, toJSON(blockchain)})
		case responseBlockchain:
			handleResponseBlockchain(n, payload.MessageText)
		}

	} else {
		log.Println(err)
	}
}
