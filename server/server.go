package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {
}

func (w *World) HelloWorld(req string, res *string) error {
	*res = req + "你好"
	return nil
}

func main() {
	// 1. 注册微服务
	err := rpc.RegisterName("hello", new(World))
	if err != nil {
		fmt.Println("微服务注册失败", err)
		return
	}
	// 2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Listen error", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听 ...")

	// 3. 建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		// go rpc.ServeConn(conn) // 只能go语言rpc，无法跨语言
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 支持跨语言

	}
}
