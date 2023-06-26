package main

import "fmt"

// 我是单行注释
func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7}
	//num = num[1:]
	//num = num[:2]
	//num = num[3:]
	//num = num[:len(num)-3] //从后面删除
	//num = num[4:] //从前面删除
	num = append(num[:len(num)-3], num[4:]...)
	//fmt.Println(num[4:])
	fmt.Println(num)
	//删除某个索引对应的值

}

/*
多行注释
*/
