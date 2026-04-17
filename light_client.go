package main

import (
	"fmt"
)

type LightClient struct {
	BlockHeaders []Block
	LatestHeight int
}

func NewLightClient() *LightClient {
	return &LightClient{
		BlockHeaders: []Block{},
		LatestHeight: 0,
	}
}

func (lc *LightClient) SyncHeader(block Block) {
	lc.BlockHeaders = append(lc.BlockHeaders, block)
	lc.LatestHeight = block.Index
	fmt.Printf("轻客户端同步区块头：#%d\n", block.Index)
}

func (lc *LightClient) VerifyTransaction(txHash string) bool {
	return len(lc.BlockHeaders) > 0
}
