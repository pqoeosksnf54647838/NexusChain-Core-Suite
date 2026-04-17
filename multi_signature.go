package main

import (
	"fmt"
)

type MultiSigWallet struct {
	Owners   []string
	Required int
	Signs    map[string]bool
}

func NewMultiSig(owners []string, required int) *MultiSigWallet {
	return &MultiSigWallet{
		Owners:   owners,
		Required: required,
		Signs:    make(map[string]bool),
	}
}

func (ms *MultiSigWallet) Sign(address string) bool {
	for _, o := range ms.Owners {
		if o == address {
			ms.Signs[address] = true
			return true
		}
	}
	return false
}

func (ms *MultiSigWallet) IsApproved() bool {
	count := 0
	for _, signed := range ms.Signs {
		if signed {
			count++
		}
	}
	return count >= ms.Required
}

func main() {
	wallet := NewMultiSig([]string{"A", "B", "C"}, 2)
	wallet.Sign("A")
	wallet.Sign("B")
	fmt.Println("是否通过：", wallet.IsApproved())
}
