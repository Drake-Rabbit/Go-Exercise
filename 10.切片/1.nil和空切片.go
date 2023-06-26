package main

import "fmt"

// 我是单行注释
func main() {

	var str []string          //	nil
	str1 := make([]string, 0) //空切片
	str2 := []string{}        //空切片
	fmt.Println(len(str), str == nil)
	fmt.Println(len(str1), str1 == nil)
	fmt.Println(len(str2), str2 == nil)

}

/*
多行注释
*/
