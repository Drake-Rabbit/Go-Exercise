package main

import (
	_ "Go-Exercise/12.go语言之函数/init1" //_下划线，即只想初始化一下，
	"fmt"
)

var test = g()

func init() {
	fmt.Println("init1")
}

func g() int {
	fmt.Println("ggg11ggg")
	return 2
}

func main() {
	//init1.Add()
	fmt.Println("main")
}

/*
多行注释
*/
