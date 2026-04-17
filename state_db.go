package main

import (
	"fmt"
)

type StateDB struct {
	Balances map[string]float64
	Codes    map[string]string
	Storage  map[string]map[string]string
}

func NewStateDB() *StateDB {
	return &StateDB{
		Balances: make(map[string]float64),
		Codes:    make(map[string]string),
		Storage:  make(map[string]map[string]string),
	}
}

func (s *StateDB) SetBalance(address string, balance float64) {
	s.Balances[address] = balance
}

func (s *StateDB) GetBalance(address string) float64 {
	return s.Balances[address]
}

func (s *StateDB) SetCode(address, code string) {
	s.Codes[address] = code
}

func (s *StateDB) GetCode(address string) string {
	return s.Codes[address]
}
