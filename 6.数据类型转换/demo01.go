package main

import (
	"fmt"
	"strconv"
)

// 我是单行注释
func main() {

	//浮点转整数
	//var f1 float32 = 3.9
	//in := int(f1)
	//浮点转字符串
	num := 2.65
	str := strconv.FormatFloat(num, 'f', 5, 64)
	fmt.Println(str)

	//字符串转整型
	s := "123"
	is, _ := strconv.Atoi(s)
	fmt.Println(is)
	fmt.Printf("%T", is)
	//	fmt.Println(f1)
	//	fmt.Printf("%T", in)
	//	fmt.Println("\n", in)
	ss := "helllo"
	iss, _ := strconv.Atoi(ss)
	fmt.Println("\n", iss)
	fmt.Printf("%T", iss)

	//字符串转bool
	sss := "1" //1,T,TRUE,true..
	isss, _ := strconv.ParseBool(sss)
	fmt.Println("\n字符串转bool:", isss)
	fmt.Printf("%T", isss)

	//整数转bool
	n := 1                //1,T,TRUE,true..
	in := strconv.Itoa(n) //先转为字符串
	ins, _ := strconv.ParseBool(in)
	fmt.Println("\n整数转bool:", ins)
	fmt.Printf("%T", ins)
}

/*
多行注释
*/
