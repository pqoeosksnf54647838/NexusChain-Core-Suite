package main

import (
	"fmt"
)

type ReorgManager struct {
	MainChain    []Block
	ForkChain    []Block
	ReorgHeight  int
}

func NewReorgManager() *ReorgManager {
	return &ReorgManager{
		MainChain:   []Block{},
		ForkChain:   []Block{},
		ReorgHeight: -1,
	}
}

func (rm *ReorgManager) DetectFork() bool {
	return len(rm.ForkChain) > len(rm.MainChain)
}

func (rm *ReorgManager) ExecuteReorg() {
	if rm.DetectFork() {
		rm.MainChain = rm.ForkChain
		rm.ReorgHeight = len(rm.MainChain)
		fmt.Println("链重组完成，切换至最长链")
	}
}
