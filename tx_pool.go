package main

import (
	"container/list"
	"fmt"
)

type TransactionPool struct {
	pool *list.List
}

func NewTxPool() *TransactionPool {
	return &TransactionPool{
		pool: list.New(),
	}
}

func (tp *TransactionPool) Add(tx string) {
	tp.pool.PushBack(tx)
	fmt.Printf("交易已加入交易池：%s\n", tx)
}

func (tp *TransactionPool) Remove(tx string) bool {
	for e := tp.pool.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == tx {
			tp.pool.Remove(e)
			return true
		}
	}
	return false
}

func (tp *TransactionPool) GetPending(count int) []string {
	var res []string
	e := tp.pool.Front()
	for i := 0; i < count && e != nil; i++ {
		res = append(res, e.Value.(string))
		e = e.Next()
	}
	return res
}

func (tp *TransactionPool) Size() int {
	return tp.pool.Len()
}
