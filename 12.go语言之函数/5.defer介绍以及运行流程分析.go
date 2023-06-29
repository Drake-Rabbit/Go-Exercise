package main

import "fmt"

func main() {
	//defer 延迟执行
	//fmt.Println("hello")
	//defer fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
	//fmt.Println("world")

	res := test1()
	fmt.Println("res", res)
	ress := test2()
	fmt.Println("ress", ress)

}
func test1() int {
	num := 5
	defer func() { //匿名函数
		num++
		fmt.Println("defer:", num)
	}()
	fmt.Println("num", num)
	return num
}

//num 5
//defer: 6
//5

func test2() (num int) {
	num = 5
	defer func() { //匿名函数
		num++
		fmt.Println("defer:", num)
	}()
	fmt.Println("num", num)
	return num
}

//num 5
//defer: 6
//res 6

/*
多行注释
*/
