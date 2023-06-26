package main

import (
	"fmt"
)

// 从P11开始
func main() {
	// := 自动推导

	name := "kuangshen"
	age := 25
	addr := "kuangshen"

	// string： 默认 空
	// int  默认 0
	fmt.Println(name, age, addr)
	fmt.Printf("%T,%T", name, addr)
}
