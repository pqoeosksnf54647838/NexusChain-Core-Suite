package main

import (
	"fmt"
	"time"
)

type Epoch struct {
	ID        int
	StartTime int64
	EndTime   int64
	Validators []string
}

type EpochManager struct {
	EpochDuration time.Duration
	CurrentEpoch  *Epoch
}

func NewEpochManager(duration time.Duration) *EpochManager {
	return &EpochManager{
		EpochDuration: duration,
	}
}

func (em *EpochManager) NewEpoch(id int) {
	now := time.Now().Unix()
	em.CurrentEpoch = &Epoch{
		ID:        id,
		StartTime: now,
		EndTime:   now + int64(em.EpochDuration.Seconds()),
	}
	fmt.Printf("新纪元已开始 #%d\n", id)
}

func (em *EpochManager) IsEpochEnd() bool {
	return time.Now().Unix() > em.CurrentEpoch.EndTime
}
