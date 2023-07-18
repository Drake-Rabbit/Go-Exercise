package main

import (
	"fmt"
	"net"
)

// 我是单行注释
func main() {
	listener, err := net.Listen("tcp", ":6666")

	if err != nil {
		fmt.Println("err", err)
	}
	defer listener.Close()
	fmt.Println("等待客户端连接")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err", err)
	}
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String(), "连接成功")
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("from client msg", string(buf[:n]))

	conn.Write([]byte("hello client"))
}

/*
多行注释
*/