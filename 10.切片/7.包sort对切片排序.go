package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{8, 3, 5, 9, 2}
	//sort.Sort(sort.Reverse(sort.IntSlice(num))) //降序

	//fmt.Println(sort.IntsAreSorted(num))
	sort.Ints(num) //Ints 以升序排列int
	//sort.Reverse(num)
	//fmt.Println(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)

	//字符串的排序
	str := []string{"f", "g", "a", "e", "be"}
	//sort.Strings(str) //默认升序

	sort.Sort(sort.Reverse(sort.StringSlice(str))) //降序
	fmt.Println(str)

}
