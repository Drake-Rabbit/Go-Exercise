package main

import "fmt"

// 我是单行注释
func main() {
	//切片append与底层数组分离
	fruit := []string{"A", "B", "C", "D", "E", "F"}
	myFruit := fruit[2:4] //当配置的容量K-i=j-i长度,append会去分离底层数组
	//长度 = j - i
	//容量 = k - i
	fmt.Printf("len=%d cap=%d %v\n", len(myFruit), cap(myFruit), myFruit)
	myFruit = append(myFruit, "1111")
	myFruit = append(myFruit, "2222")
	fmt.Println(fruit, myFruit)

}

/*
多行注释
*/
