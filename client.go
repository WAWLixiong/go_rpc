package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 用 rpc 连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	// 调用远程函数
	var reply string
	err = conn.Call("hello.HelloWorld", "我是客户端", &reply)
	if err != nil {
		return
	}
	fmt.Println(reply)
}
