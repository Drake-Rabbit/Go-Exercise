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
func RetSum(a ...int) int {
	fmt.Printf("%T,%v", a, a)
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum

}

func main() {

	res := RetSum(3, 4, 5, 6, 7)
	fmt.Println(res)
}

/*
多行注释
*/
