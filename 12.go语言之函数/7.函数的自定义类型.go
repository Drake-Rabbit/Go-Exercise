package main

import "fmt"

// 自定义类型
type myInt int
type myString = string

// 自定义函数类型
type myType func(int, int) int

func adw(a, b int) int {
	return a + b
}

func main() {
	var myFun myType
	myFun = adw
	fmt.Printf("%T\n", myFun)
	fmt.Printf("%T\n", adw)

	//var num myInt = 25
	//var str myString = "hello world"
	//fmt.Println(num)
	//fmt.Printf("%T\n", num)
	//fmt.Println(str)
	//fmt.Printf("%T\n", str)
	//
	//var n rune = 23
	//fmt.Println(n)
	//fmt.Printf("%T\n", n)
	//
	//var s rune = 'h'
	//fmt.Println(s)
	//fmt.Printf("%T\n", s)
	//
	//var i byte = 7
	//fmt.Println(i)
	//fmt.Printf("%T\n", i)
}

/*
多行注释
*/
