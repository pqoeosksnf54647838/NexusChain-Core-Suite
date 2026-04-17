package main

import (
	"fmt"
	"time"
)

type BlockBuilder struct {
	Height    int
	PrevHash  string
	Txs       []string
	Difficulty int
}

func NewBlockBuilder(height int, prevHash string, diff int) *BlockBuilder {
	return &BlockBuilder{
		Height:    height,
		PrevHash:  prevHash,
		Difficulty: diff,
		Txs:       []string{},
	}
}

func (bb *BlockBuilder) AddTransaction(tx string) {
	bb.Txs = append(bb.Txs, tx)
}

func (bb *BlockBuilder) Build() Block {
	block := Block{
		Index:      bb.Height,
		Timestamp:  time.Now().String(),
		PrevHash:   bb.PrevHash,
		Difficulty: bb.Difficulty,
	}
	for _, tx := range bb.Txs {
		block.Data += tx + ";"
	}
	block.Hash = calculateHash(block)
	fmt.Printf("区块构建完成 #%d\n", bb.Height)
	return block
}
