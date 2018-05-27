package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func calculateHash(index uint64, prevHash string, data Transaction, timestamp uint64) string {
	txData := fmt.Sprint(data.Value) + data.Input + data.Output
	payload := fmt.Sprint(index) + prevHash + txData + fmt.Sprint(timestamp)
	h := sha256.New()
	h.Write([]byte(payload))

	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateHashForBlock(block *Block) string {
	return calculateHash(block.Index, block.PreviousHash, block.Data, block.Timestamp)
}

func filterEndpointsFromConnections(connections []*websocket.Conn) (endpoints []string) {
	endpoints = make([]string, len(connections))
	for i, p := range connections {
		endpoints[i] = p.RemoteAddr().String()
	}
	return endpoints
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
