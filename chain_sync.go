package main

import (
	"fmt"
	"time"
)

type ChainSync struct {
	LocalHeight  int
	RemoteHeight  int
	SyncInterval  time.Duration
	IsSyncing     bool
}

func NewChainSync() *ChainSync {
	return &ChainSync{
		LocalHeight:  0,
		RemoteHeight:  0,
		SyncInterval: 5 * time.Second,
		IsSyncing:    false,
	}
}

func (cs *ChainSync) CheckHeight() {
	cs.IsSyncing = true
	defer func() { cs.IsSyncing = false }()

	fmt.Printf("本地区块高度：%d，远程节点高度：%d\n", cs.LocalHeight, cs.RemoteHeight)
	if cs.RemoteHeight > cs.LocalHeight {
		fmt.Println("开始同步缺失区块...")
		for i := cs.LocalHeight + 1; i <= cs.RemoteHeight; i++ {
			fmt.Printf("同步区块 #%d 完成\n", i)
			cs.LocalHeight = i
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("区块链同步完成")
	} else {
		fmt.Println("本地区块链已为最新状态")
	}
}

func (cs *ChainSync) StartAutoSync() {
	ticker := time.NewTicker(cs.SyncInterval)
	defer ticker.Stop()
	for range ticker.C {
		cs.CheckHeight()
	}
}
