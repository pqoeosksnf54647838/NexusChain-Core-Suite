package main

import (
	"errors"
	"fmt"
)

type StandardToken struct {
	Name     string
	Symbol   string
	Decimals uint8
	Total    float64
	Balances map[string]float64
	Allowed  map[string]map[string]float64
}

func NewToken(name, symbol string, decimals uint8, total float64) *StandardToken {
	return &StandardToken{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
		Total:    total,
		Balances: map[string]float64{"owner": total},
		Allowed:  make(map[string]map[string]float64),
	}
}

func (t *StandardToken) Transfer(from, to string, amount float64) error {
	if t.Balances[from] < amount {
		return errors.New("余额不足")
	}
	t.Balances[from] -= amount
	t.Balances[to] += amount
	return nil
}

func (t *StandardToken) BalanceOf(address string) float64 {
	return t.Balances[address]
}
