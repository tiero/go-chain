package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

//MessageType struct
type MessageType struct {
	queryLatestBlock   int
	queryAllBlock      int
	responseBlockchain int
}

var messageType = MessageType{0, 1, 2}

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

		mt, message, err1 := c.ReadMessage()

		if err == nil {
			//c.WriteMessage(0, []byte("ciao"))
			log.Println(mt)
			log.Println(err1)
			err2 := c.WriteMessage(mt, message)
			if err2 != nil {
				log.Println(err2)
			} else {
				log.Printf("evviva")
				node.Peers = append(node.Peers, c)
			}
		} else {
			log.Println(err)
		}
	}
}

func handleIncomingMessage(conn *websocket.Conn, msgType int, message []byte) {

	err := conn.WriteMessage(msgType, message)
	if err != nil {
		log.Println("write:", err)
	} else {
		log.Println("sent")
	}

	/* fmt.Println(message)
	switch msgType {
	case messageType.queryLatestBlock:

	} */

}
