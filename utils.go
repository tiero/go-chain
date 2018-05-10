package main

import (
	"fmt"
	"encoding/hex"
	"crypto/sha256"	
)

func calculateHash(index int, prevHash string, data Transaction, timestamp int64) string {
	payload := fmt.Sprint(index) + prevHash + fmt.Sprint(data.value) + data.input + data.output + fmt.Sprint(timestamp)
	h := sha256.New()
	h.Write([]byte(payload))

	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)	
}

func genesisBlock() Block {
	// Wednesday 9th May 2018 10:16:19 PM UTC
	return Block{0, "3cd45a480c2601ed55245eac8b233c680f111eaad30c568a318e5213f7f0f522", "0", Transaction{}, 1525904179}
}
