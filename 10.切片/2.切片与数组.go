package main

import "fmt"

// 我是单行注释
func main() {

	num := [6]int{1, 2, 3, 4, 5, 6}
	s := num[1:4] //左闭右开  slice[i:j:k] k控制切片的容量==k-i<=原来的长度
	//长度 = j - i
	//容量 = k - i
	print(s, "\n")
	fmt.Println(num)
	fmt.Println(s)

	//切片共享同一个底层数组
	mynum := []int{10, 20, 30, 40, 50}
	newNum := mynum[1:3]
	mynum[1] = 100
	fmt.Println(mynum, newNum)

}

/*
多行注释
*/
