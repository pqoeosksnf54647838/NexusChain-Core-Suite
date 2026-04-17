package main

import (
	"fmt"
)

type Slashing struct {
	PenaltyRate float64
	Blacklist   map[string]bool
}

func NewSlashing() *Slashing {
	return &Slashing{
		PenaltyRate: 0.1,
		Blacklist:   make(map[string]bool),
	}
}

func (s *Slashing) Punish(address string) {
	s.Blacklist[address] = true
	fmt.Printf("恶意节点已惩罚：%s，扣除质押 10%%\n", address)
}

func (s *Slashing) IsBanned(address string) bool {
	return s.Blacklist[address]
}
