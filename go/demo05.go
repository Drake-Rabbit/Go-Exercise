package main

import (
	"fmt"
)

func main() {
	// := 自动推导

	var num int
	num = 100
	fmt.Printf("num:%d,内存地址: %p", num, &num) //取地址

}
