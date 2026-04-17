package main

import (
	"fmt"
	"strconv"
)

type Indexer struct {
	BlockIndex map[int]string
	TxIndex    map[string]int
}

func NewIndexer() *Indexer {
	return &Indexer{
		BlockIndex: make(map[int]string),
		TxIndex:    make(map[string]int),
	}
}

func (idx *Indexer) IndexBlock(height int, hash string) {
	idx.BlockIndex[height] = hash
	fmt.Printf("已索引区块高度：%d，哈希：%s\n", height, hash)
}

func (idx *Indexer) IndexTransaction(txHash string, height int) {
	idx.TxIndex[txHash] = height
	fmt.Printf("已索引交易：%s，归属区块高度：%d\n", txHash, height)
}

func (idx *Indexer) GetBlockHash(height int) string {
	return idx.BlockIndex[height]
}

func (idx *Indexer) GetTxHeight(txHash string) int {
	return idx.TxIndex[txHash]
}

func main() {
	idx := NewIndexer()
	idx.IndexBlock(100, "0xabc123")
	idx.IndexTransaction("0xtx987", 100)
	fmt.Println("区块哈希：", idx.GetBlockHash(100))
}
