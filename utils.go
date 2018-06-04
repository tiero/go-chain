package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func calculateHash(index uint64, term int, prevHash string, data Transaction, timestamp uint64) string {
	txData := fmt.Sprint(data.Value) + data.Input + data.Output
	payload := fmt.Sprint(index) + fmt.Sprint(term) + prevHash + txData + fmt.Sprint(timestamp)
	h := sha256.New()
	h.Write([]byte(payload))

	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateHashForBlock(block *Block) string {
	return calculateHash(block.Index, block.Term, block.PreviousHash, block.Data, block.Timestamp)
}

func logState(n *Node) string {
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

func filterEndpointsFromConnections(connections map[string]*websocket.Conn) (endpoints []string) {
	keys := make([]string, len(connections))

	i := 0
	for k := range connections {
		log.Println("∆========")
		println(k)
		println(connections[k].RemoteAddr().String())
		log.Println("========")
		keys[i] = k
		i++
	}

	return keys
}

func filterConnectionFromEndpoint(connections []*websocket.Conn, endpoint string) *websocket.Conn {
	conn := &websocket.Conn{}
	for _, c := range connections {
		log.Println("========")
		log.Println(c.RemoteAddr().String())
		log.Println(endpoint)
		log.Println("========")
		if c.RemoteAddr().String() == endpoint {
			conn = c
		}
	}
	return conn
}

func setTimeout(duration time.Duration) <-chan time.Time {
	if duration == 0 {
		return nil
	}
	return time.After(duration + (5 * time.Millisecond))
}
