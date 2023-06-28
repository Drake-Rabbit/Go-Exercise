package main

import (
	"fmt"
)

// 我是单行注释
/**
* @Description:
* @Author: oDuck
* @Date: 2023/6/28
 */

// 可变参数
func ssssetSum(a int, b ...int) int {
	println()
	fmt.Println(a, b)
	return 5
}

func RetSum(a ...int) int {
	fmt.Printf("%T,%v", a, a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

// 返回值
// func get(a,b int)int
func get(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub

}

// 返回值的命名  相当于在参数里提前声明了
func col(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub

}

func main() {

	res := RetSum(3, 4, 5, 6, 7)
	fmt.Println(res)
	rrr := ssssetSum(1, 2, 3, 4, 5)
	fmt.Println(rrr)

	hd, sd := get(8, 2)
	fmt.Println(hd, sd)
}

/*
多行注释
*/
