package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
}

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d%d", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce, block.Difficulty)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock Block, data string, difficulty int) Block {
	var newBlock Block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Difficulty = difficulty
	newBlock.Nonce = 0

	for {
		hash := calculateHash(newBlock)
		if hash[:difficulty] == repeatZero(difficulty) {
			newBlock.Hash = hash
			break
		}
		newBlock.Nonce++
	}
	return newBlock
}

func repeatZero(n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += "0"
	}
	return res
}

func main() {
	genesisBlock := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "genesis block",
		PrevHash:   "0",
		Nonce:      0,
		Difficulty: 2,
	}
	genesisBlock.Hash = calculateHash(genesisBlock)
	fmt.Printf("创世区块：%+v\n", genesisBlock)
}
