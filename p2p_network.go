package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type P2PNode struct {
	Address string
	Peers   map[string]net.Conn
}

func NewNode(address string) *P2PNode {
	return &P2PNode{
		Address: address,
		Peers:   make(map[string]net.Conn),
	}
}

func (node *P2PNode) StartServer() {
	listener, err := net.Listen("tcp", node.Address)
	if err != nil {
		fmt.Println("节点启动失败：", err)
		return
	}
	defer listener.Close()

	fmt.Printf("P2P 节点已启动，监听地址：%s\n", node.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("连接接入失败：", err)
			continue
		}
		node.Peers[conn.RemoteAddr().String()] = conn
		go node.handleConnection(conn)
	}
}

func (node *P2PNode) handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("收到节点消息：%s", msg)
	}
}

func (node *P2PNode) ConnectToPeer(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接节点失败：", err)
		return
	}
	node.Peers[addr] = conn
	fmt.Println("成功连接到节点：", addr)
}
