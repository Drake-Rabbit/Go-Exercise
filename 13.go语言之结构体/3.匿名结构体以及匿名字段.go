package main

import "fmt"

// 匿名字段
type P struct {
	string
	int
}

func main() {
	//匿名结构体
	var p struct {
		name string
		age  int
	}
	p.name = "lam"
	p.age = 369
	fmt.Printf("%#v,%T", p, p)

	//匿名字段的使用
	ppp := P{
		"dasdas",
		55,
	}
	fmt.Println()
	fmt.Printf("%#v,%T", ppp, ppp)
}

/*
多行注释
*/
