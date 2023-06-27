package main

import (
	"fmt"
	"sort"
)

// 我是单行注释
func main() {

	users := make(map[int]string, 20)
	users[1] = "lampol"
	users[2] = "bbbb"
	users[3] = "dasdasda"
	users[4] = "wwwwwwa"
	users[5] = "dddddda"
	fmt.Println("map----:", users)

	//判断 键值对 是否存在
	//返回两个值，第一个是key对应的value,第二个是bool；存在返回true
	_, isEx := users[1]
	fmt.Println(isEx)

	fmt.Println("遍历map:")
	//遍历 map无序,只能range遍历
	for k, v := range users {
		fmt.Println(k, v)
	}

	//对map进行升序遍历,key
	keys := make([]int, 0, len(users))
	for k, _ := range users {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	fmt.Println("打印升序的map:")
	for _, v := range keys {
		fmt.Println(v, users[v])
	}

}
