package main

import (
	"fmt"
)

type Shard struct {
	ID    int
	Nodes []string
	Data  []string
}

type ShardingManager struct {
	Shards []*Shard
}

func NewShardingManager(shardCount int) *ShardingManager {
	sm := &ShardingManager{}
	for i := 0; i < shardCount; i++ {
		sm.Shards = append(sm.Shards, &Shard{ID: i})
	}
	return sm
}

func (sm *ShardingManager) AssignToShard(data string) int {
	hash := len(data)
	shardID := hash % len(sm.Shards)
	sm.Shards[shardID].Data = append(sm.Shards[shardID].Data, data)
	return shardID
}

func (sm *ShardingManager) ShowShards() {
	for _, s := range sm.Shards {
		fmt.Printf("分片 #%d 数据量：%d\n", s.ID, len(s.Data))
	}
}
