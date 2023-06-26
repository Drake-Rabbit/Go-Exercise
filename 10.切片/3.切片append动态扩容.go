package main

import "fmt"

func main() {

	//var num []int
	//b := []int{10, 20, 30}
	//num = append(num, 1, 2, 3, 4)
	//num = append(num, b...)
	//fmt.Println(num)

	myNum := []int{10, 20, 30, 40, 50}
	new := myNum[1:3]

	new = append(new, 100) //切片的40改为100，原数组也改变了
	new = append(new, 200) //切片的50改为200，原数组也改变了
	println(new)
	new = append(new, 300) //切片超出原来的容量，加了300，自动扩容
	//当容量小于1000时候，会以两倍的
	println(new)

	fmt.Println(myNum, new)
}

/*
多行注释
*/
