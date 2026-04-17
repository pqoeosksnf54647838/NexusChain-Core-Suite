package main

import (
	"fmt"
	"sort"
)

type FeeOrder struct {
	TxID string
	Fee  float64
}

type FeeMarket struct {
	Orders []FeeOrder
}

func NewFeeMarket() *FeeMarket {
	return &FeeMarket{}
}

func (fm *FeeMarket) AddOrder(txID string, fee float64) {
	fm.Orders = append(fm.Orders, FeeOrder{TxID: txID, Fee: fee})
}

func (fm *FeeMarket) SortByFee() {
	sort.Slice(fm.Orders, func(i, j int) bool {
		return fm.Orders[i].Fee > fm.Orders[j].Fee
	})
}

func (fm *FeeMarket) ShowTop5() {
	count := 5
	if len(fm.Orders) < 5 {
		count = len(fm.Orders)
	}
	for i := 0; i < count; i++ {
		fmt.Printf("交易 %s 手续费 %.4f\n", fm.Orders[i].TxID, fm.Orders[i].Fee)
	}
}
