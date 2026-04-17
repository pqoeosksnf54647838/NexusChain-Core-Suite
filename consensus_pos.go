package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   int
	Active  bool
}

type POSConsensus struct {
	Validators []Validator
}

func NewPOS() *POSConsensus {
	return &POSConsensus{
		Validators: []Validator{},
	}
}

func (pos *POSConsensus) RegisterValidator(address string, stake int) {
	pos.Validators = append(pos.Validators, Validator{
		Address: address,
		Stake:   stake,
		Active:  true,
	})
}

func (pos *POSConsensus) SelectValidator() *Validator {
	var activeValidators []Validator
	for _, v := range pos.Validators {
		if v.Active {
			activeValidators = append(activeValidators, v)
		}
	}

	if len(activeValidators) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	totalWeight := 0
	for _, v := range activeValidators {
		totalWeight += v.Stake
	}
	target := rand.Intn(totalWeight)
	current := 0
	for _, v := range activeValidators {
		current += v.Stake
		if current > target {
			return &v
		}
	}
	return &activeValidators[0]
}

func main() {
	pos := NewPOS()
	pos.RegisterValidator("node1", 100)
	pos.RegisterValidator("node2", 300)
	pos.RegisterValidator("node3", 200)

	selected := pos.SelectValidator()
	fmt.Printf("当选区块验证节点：%s，质押金额：%d\n", selected.Address, selected.Stake)
}
