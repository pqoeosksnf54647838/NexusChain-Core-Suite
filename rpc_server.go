package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ChainRPC int

type RPCArgs struct {
	Height int
	Tx     string
}

type RPCReply struct {
	Hash  string
	Value string
}

func (c *ChainRPC) GetBlockHash(args *RPCArgs, reply *RPCReply) error {
	reply.Hash = fmt.Sprintf("hash_for_height_%d", args.Height)
	return nil
}

func StartRPCServer() {
	chain := new(ChainRPC)
	rpc.Register(chain)
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("RPC 启动失败：", err)
		return
	}
	fmt.Println("RPC 服务已启动：:8888")
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
