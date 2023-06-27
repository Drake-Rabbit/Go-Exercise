package main

import "fmt"

// 我是单行注释
func main() {

	users := map[int]string{1: "a", 2: "b"}
	fmt.Println(users)
	fmt.Printf("%T", users)
	fmt.Println()

	//var uses  map[int]string
	//还需要分配 空间
	var u map[int]string
	//u =make(map[int]string,4) //容量可以不固定，因为可以动态扩容
	u = make(map[int]string)
	u[1] = "aa"
	u[2] = "bb"
	fmt.Println(u, u[1])

	fmt.Println("删除前:", u[1])
	//删除
	delete(u, 1)
	fmt.Println("删除后:", u[1])

}

/*
多行注释
*/
