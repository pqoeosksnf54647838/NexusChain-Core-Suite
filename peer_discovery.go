package main

import (
	"fmt"
	"math/rand"
	"time"
)

type PeerDiscovery struct {
	BootstrapNodes []string
	ActivePeers    []string
}

func NewPeerDiscovery() *PeerDiscovery {
	return &PeerDiscovery{
		BootstrapNodes: []string{
			"node1.chain.net:9000",
			"node2.chain.net:9000",
			"node3.chain.net:9000",
		},
		ActivePeers: []string{},
	}
}

func (pd *PeerDiscovery) Discover() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("开始节点发现...")
	for _, node := range pd.BootstrapNodes {
		if rand.Intn(2) == 1 {
			pd.ActivePeers = append(pd.ActivePeers, node)
			fmt.Printf("发现可用节点：%s\n", node)
		}
	}
	fmt.Printf("节点发现完成，可用节点数：%d\n", len(pd.ActivePeers))
}

func (pd *PeerDiscovery) GetRandomPeer() string {
	if len(pd.ActivePeers) == 0 {
		return ""
	}
	return pd.ActivePeers[rand.Intn(len(pd.ActivePeers))]
}
