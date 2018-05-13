package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

// Node represent the current operating daemon
type Node struct {
	Host   string
	Peers  []string
	Height uint64
}

// NewNode starts a p2p server
func NewNode(host string, initialPeers []string, blocksHeight uint64) *Node {
	mutex.Lock()
	defer mutex.Unlock()
	return &Node{host, initialPeers, blocksHeight}
}

func connectToPeers(node *Node, endpoints []string) {
	for _, endpoint := range endpoints {
		parsedURL, _ := url.Parse(endpoint)
		u := url.URL{Scheme: parsedURL.Scheme, Host: parsedURL.Host, Path: "/ws"}
		log.Printf("connecting to %s", u.String())

		_, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

		if err == nil {
			node.Peers = append(node.Peers, endpoint)
		} else {
			log.Println(err)
		}
	}
}
