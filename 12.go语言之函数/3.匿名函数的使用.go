package main

import "fmt"

// 利用匿名函数实现回调遍历
func forSlic(s []string, f func(i string)) {
	for _, v := range s {
		f(v)
	}
}

// 匿名函数
func main() {

	sli := []string{"a", "b", "c"}
	forSlic(sli, func(v string) {
		fmt.Println(v)
	})

	func(str string) {
		fmt.Println(str)
	}("hello world")

	sum := func(a, b int) int {
		return a + b
	}
	res := sum(5, 6)
	fmt.Println(res)
}

/*
多行注释
*/
