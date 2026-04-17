package main

import (
	"fmt"
	"math"
)

type RewardPool struct {
	TotalReward float64
	APR         float64
	Stakers     map[string]float64
}

func NewRewardPool(total float64, apr float64) *RewardPool {
	return &RewardPool{
		TotalReward: total,
		APR:         apr,
		Stakers:     make(map[string]float64),
	}
}

func (rp *RewardPool) Stake(address string, amount float64) {
	rp.Stakers[address] = amount
}

func (rp *RewardPool) CalculateReward(address string, days int) float64 {
	stake := rp.Stakers[address]
	rate := rp.APR / 365
	return stake * rate * float64(days)
}

func (rp *RewardPool) Distribute() {
	for addr, stake := range rp.Stakers {
		reward := rp.CalculateReward(addr, 30)
		fmt.Printf("地址 %s 30天质押奖励：%.4f\n", addr, reward)
	}
}
