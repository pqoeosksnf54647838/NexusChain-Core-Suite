package main

import (
	"fmt"
)

type UTXO struct {
	TxHash  string
	Index   int
	Address string
	Amount  float64
	Spent   bool
}

type UTXOSet struct {
	UTXOs map[string]UTXO
}

func NewUTXOSet() *UTXOSet {
	return &UTXOSet{
		UTXOs: make(map[string]UTXO),
	}
}

func (set *UTXOSet) AddUTXO(txHash string, index int, address string, amount float64) {
	key := fmt.Sprintf("%s_%d", txHash, index)
	set.UTXOs[key] = UTXO{
		TxHash:  txHash,
		Index:   index,
		Address: address,
		Amount:  amount,
		Spent:   false,
	}
}

func (set *UTXOSet) SpendUTXO(key string) {
	if utxo, exists := set.UTXOs[key]; exists {
		utxo.Spent = true
		set.UTXOs[key] = utxo
	}
}

func (set *UTXOSet) GetBalance(address string) float64 {
	balance := 0.0
	for _, utxo := range set.UTXOs {
		if utxo.Address == address && !utxo.Spent {
			balance += utxo.Amount
		}
	}
	return balance
}
