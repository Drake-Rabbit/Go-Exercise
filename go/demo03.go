package main

import "fmt"

func main() {
	//
	var (
		name string
		age  int
		addr string
	)

	name = "kuangshen"
	age = 25
	addr = "kuangshen"

	// string： 默认 空
	// int  默认 0
	fmt.Println(name, age, addr)
}
