package main

import (
	"errors"
	"fmt"
)

type VM struct {
	Storage map[string]string
	Stack   []int
}

func NewVM() *VM {
	return &VM{
		Storage: make(map[string]string),
		Stack:   []int{},
	}
}

func (vm *VM) Push(val int) {
	vm.Stack = append(vm.Stack, val)
}

func (vm *VM) Pop() (int, error) {
	if len(vm.Stack) == 0 {
		return 0, errors.New("栈为空")
	}
	val := vm.Stack[len(vm.Stack)-1]
	vm.Stack = vm.Stack[:len(vm.Stack)-1]
	return val, nil
}

func (vm *VM) Add() error {
	a, err := vm.Pop()
	if err != nil {
		return err
	}
	b, err := vm.Pop()
	if err != nil {
		return err
	}
	vm.Push(a + b)
	return nil
}

func (vm *VM) Set(key, value string) {
	vm.Storage[key] = value
}

func (vm *VM) Get(key string) string {
	return vm.Storage[key]
}

func main() {
	vm := NewVM()
	vm.Push(10)
	vm.Push(20)
	vm.Add()
	res, _ := vm.Pop()
	fmt.Println("计算结果：", res)
}
