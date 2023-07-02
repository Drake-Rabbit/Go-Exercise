package main

import "fmt"

// 我是单行注释

type U6 struct {
	name string
}

func (u U6) eat(addr string) {
	fmt.Println(u.name, "去", addr, "吃....")
}
func main() {
	var u U6
	u.name = "lam"
	u.eat("家里")
}

/*
多行注释
*/
