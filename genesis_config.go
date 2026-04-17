package main

import (
	"encoding/json"
	"fmt"
)

type GenesisConfig struct {
	ChainName    string  `json:"chain_name"`
	InitialSupply float64 `json:"initial_supply"`
	BlockTime    int     `json:"block_time"`
	Difficulty   int     `json:"difficulty"`
}

func LoadGenesis() GenesisConfig {
	raw := `{
		"chain_name": "nexus-chain",
		"initial_supply": 100000000,
		"block_time": 3,
		"difficulty": 2
	}`
	var config GenesisConfig
	json.Unmarshal([]byte(raw), &config)
	return config
}

func main() {
	genesis := LoadGenesis()
	fmt.Printf("创世配置：%+v\n", genesis)
}
