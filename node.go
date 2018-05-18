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

//MessagePayload struct
type MessagePayload struct {
	MessageType int
	MessageText string
}

// Node represent the current operating daemon
type Node struct {
	Host   string
	Peers  []*websocket.Conn
	Height uint64
}

// NewNode starts a p2p server
func NewNode(host string, initialPeers []*websocket.Conn, blocksHeight uint64) *Node {
	mutex.Lock()
	defer mutex.Unlock()
	return &Node{host, initialPeers, blocksHeight}
}

func connectToPeers(node *Node, endpoints []string) []string {
	var connected []string
	for _, endpoint := range endpoints {
		//Parse endpoints
		// TODO move to helper function
		parsedURL, _ := url.Parse(endpoint)
		u := url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host, Path: "/ws"}
		//Start a dialing
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

		if err == nil {
			go wsListen(node, c)
			connected = append(connected, endpoint)
			//c.WriteMessage(websocket.TextMessage, []byte("ciao cazzu"))
			broadcastMessage(c, &MessagePayload{queryLatestBlock, ""})
		} else {
			log.Println(err)
		}
	}
	return connected
}

func wsListen(node *Node, conn *websocket.Conn) {
	node.Peers = append(node.Peers, conn)
	println(conn.RemoteAddr().String() + " Connected")
	println("Peers: " + fmt.Sprint(len(node.Peers)))
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			println(conn.RemoteAddr().String() + " Disconnected")
			removePeer(node, conn)
			println("Peers: " + fmt.Sprint(len(node.Peers)))
			return
		}
		handleIncomingMessage(conn, p)
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

func broadcastMessage(conn *websocket.Conn, mp *MessagePayload) {
	payload, err := json.Marshal(mp)
	if err == nil {
		if err = conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			return
		}
	}

}

func handleResponseBlockchain(message string) {
	receivedBlockchain := fromJSON(message)
	println("reveived blocks lenght " + fmt.Sprint(len(receivedBlockchain.blocks)))
}

func handleIncomingMessage(conn *websocket.Conn, message []byte) {
	var payload MessagePayload
	if err := json.Unmarshal(message, &payload); err == nil {
		switch int(payload.MessageType) {
		case queryLatestBlock:
			//latest block
			broadcastMessage(conn, &MessagePayload{responseBlockchain, toJSON(blockchain, true)})
		case queryAllBlock:
			broadcastMessage(conn, &MessagePayload{responseBlockchain, toJSON(blockchain)})
		case responseBlockchain:
			handleResponseBlockchain(payload.MessageText)
		}

	} else {
		log.Println(err)
	}
}
