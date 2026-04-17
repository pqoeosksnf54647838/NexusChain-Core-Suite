package main

import (
	"fmt"
)

type ContractStorage struct {
	Data map[string]string
}

func NewContractStorage() *ContractStorage {
	return &ContractStorage{
		Data: make(map[string]string),
	}
}

func (cs *ContractStorage) Set(key, value string) {
	cs.Data[key] = value
	fmt.Printf("合约存储：%s = %s\n", key, value)
}

func (cs *ContractStorage) Get(key string) string {
	return cs.Data[key]
}

func (cs *ContractStorage) Delete(key string) {
	delete(cs.Data, key)
	fmt.Printf("删除存储键：%s\n", key)
}

func (cs *ContractStorage) Exists(key string) bool {
	_, ok := cs.Data[key]
	return ok
}
