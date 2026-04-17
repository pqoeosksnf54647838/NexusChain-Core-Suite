package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartCLI() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("区块链节点 CLI 已启动")
	fmt.Println("支持命令：createblock, peers, balance, exit")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "createblock":
			fmt.Println("新区块已生成")
		case "peers":
			fmt.Println("在线节点数：8")
		case "balance":
			fmt.Println("账户余额：100.0")
		case "exit":
			fmt.Println("退出 CLI")
			return
		default:
			fmt.Println("未知命令")
		}
	}
}

func main() {
	StartCLI()
}
