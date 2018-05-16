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

func connectToPeers(node *Node, endpoints []string) {
	for _, endpoint := range endpoints {
		parsedURL, _ := url.Parse(endpoint)
		u := url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host, Path: "/ws"}
		log.Printf("connecting to %s", u.String())

		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

		//mt, message, err1 := c.ReadMessage()

		if err == nil {
			node.Peers = append(node.Peers, c)
			broadcastMessage(c, &MessagePayload{queryLatestBlock, ""})
		} else {
			log.Println(err)
		}
	}
}

func broadcastMessage(conn *websocket.Conn, mp *MessagePayload) {
	log.Println("siamo qua")
	payload, err := json.Marshal(mp)
	if err == nil {
		if err = conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			return
		}
	}

}

func handleIncomingMessage(conn *websocket.Conn, message []byte) {
	log.Println(message)
	var payload MessagePayload
	if err := json.Unmarshal(message, &payload); err == nil {
		log.Println(payload.MessageType)
		log.Println(payload.MessageText)
		broadcastMessage(conn, &MessagePayload{responseBlockchain, fmt.Sprint(latestBlock(blockchain).Index)})
		/* switch payload.MessageType {
		case 0:

		} */
	} else {
		log.Println(err)
	}
}
