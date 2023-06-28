package main

import "fmt"

// 我是单行注释
/**
* @Description:
* @Author: oDuck
* @Date: 2023/6/28
 */
func getSum(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(getSum(3, 4))
}

/*
多行注释
*/
