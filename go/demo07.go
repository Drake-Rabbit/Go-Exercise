package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {
	// 对象: User
	a, _ := test() // "_"匿名字符
	fmt.Println(a)
}
