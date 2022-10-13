package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 用 rpc 连接服务器
	// conn, err := rpc.Dial("tcp", "127.0.0.1:8080") // 不能跨语言
	conn, err := net.Dial("tcp", "127.0.0.1:8080") // 不能跨语言
	if err != nil {
		return
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	// 调用远程函数
	var reply string
	// err = conn.Call("hello.HelloWorld", "我是客户端", &reply) // 不能跨语言
	err = client.Call("hello.HelloWorld", "我是客户端", &reply)
	if err != nil {
		return
	}
	fmt.Println(reply)
}
