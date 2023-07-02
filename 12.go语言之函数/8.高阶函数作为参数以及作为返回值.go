package main

import "fmt"

type myType1 func(int) int

// 把函数变为返回值
func ADD() myType1 {
	i := 20
	f := func(num int) int {
		i += num
		return i
	}
	return f
}

//	func ADD() func(int) int {
//		i := 20
//		f := func(num int) int {
//			i += num
//			return i
//		}
//		return f
//	}
//
// 偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

// 奇数
func isOdd(num int) bool {
	if num%2 == 1 {
		return true
	}
	return false
}

// 把函数变为参数
type functype func(int) bool

func doNum(num []int, functype2 functype) []int {
	res := make([]int, 0, 20)
	for _, v := range num {
		if functype2(v) {
			res = append(res, v)
		}
	}
	return res
}
func main() {

	num := []int{2, 4, 5, 7, 8, 9}
	odd := doNum(num, isOdd)
	fmt.Println(odd)
	even := doNum(num, isEven)
	fmt.Println(even)

	//f := ADD()
	//num := f(5)
	//fmt.Println(num)

}

/*
多行注释
*/
