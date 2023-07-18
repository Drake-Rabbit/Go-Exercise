package main

import (
	"fmt"
	"net"
)

func main() {
	//新打开一个窗口
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("hello server"))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("from server msg", string(buf[:n]))
}
