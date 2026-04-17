package main

import (
	"fmt"
	"time"
)

type Bridge struct {
	SourceChain      string
	TargetChain      string
	PendingTransfers map[string]BridgeTransfer
}

type BridgeTransfer struct {
	ID        string
	From      string
	To        string
	Amount    float64
	Status    string
	Timestamp int64
}

func NewBridge(source, target string) *Bridge {
	return &Bridge{
		SourceChain:      source,
		TargetChain:      target,
		PendingTransfers: make(map[string]BridgeTransfer),
	}
}

func (b *Bridge) CreateTransfer(id, from, to string, amount float64) {
	b.PendingTransfers[id] = BridgeTransfer{
		ID:        id,
		From:      from,
		To:        to,
		Amount:    amount,
		Status:    "pending",
		Timestamp: time.Now().Unix(),
	}
	fmt.Printf("跨链转账已创建：%s\n", id)
}

func (b *Bridge) ConfirmTransfer(id string) {
	if t, ok := b.PendingTransfers[id]; ok {
		t.Status = "confirmed"
		b.PendingTransfers[id] = t
		fmt.Printf("跨链转账已确认：%s\n", id)
	}
}
