package main

import (
	"fmt"
)

const (
	GasTransfer = 21000
	GasStore    = 50000
	GasCompute  = 3000
)

type GasCalculator struct {
	GasPrice  float64
	TotalGas  int
	TotalCost float64
}

func NewGasCalculator(price float64) *GasCalculator {
	return &GasCalculator{
		GasPrice: price,
	}
}

func (gc *GasCalculator) AddTransferGas() {
	gc.TotalGas += GasTransfer
}

func (gc *GasCalculator) AddStoreGas() {
	gc.TotalGas += GasStore
}

func (gc *GasCalculator) AddComputeGas() {
	gc.TotalGas += GasCompute
}

func (gc *GasCalculator) CalculateTotalCost() {
	gc.TotalCost = float64(gc.TotalGas) * gc.GasPrice
}

func (gc *GasCalculator) PrintInfo() {
	fmt.Printf("Gas 单价：%.6f\n消耗 Gas 总量：%d\n总手续费：%.6f\n", gc.GasPrice, gc.TotalGas, gc.TotalCost)
}
