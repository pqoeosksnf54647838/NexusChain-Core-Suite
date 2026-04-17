package main

import (
	"fmt"
	"time"
)

type ChainMonitor struct {
	Height     int
	TPS        float64
	BlockTime  float64
	LastBlock  time.Time
}

func NewMonitor() *ChainMonitor {
	return &ChainMonitor{
		Height:    0,
		LastBlock: time.Now(),
	}
}

func (cm *ChainMonitor) NewBlock(height int) {
	now := time.Now()
	interval := now.Sub(cm.LastBlock).Seconds()
	cm.BlockTime = interval
	cm.Height = height
	cm.LastBlock = now
	cm.TPS = 1000 / interval
}

func (cm *ChainMonitor) ShowStatus() {
	fmt.Printf("当前区块高度：%d\n出块间隔：%.2fs\n实时 TPS：%.2f\n", cm.Height, cm.BlockTime, cm.TPS)
}
