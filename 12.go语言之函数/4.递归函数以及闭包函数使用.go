package main

import "fmt"

// 阶乘
func jc(num int) int {
	if num <= 1 {
		return 1
	}
	return num * jc(num-1)
}

// 闭包函数，返回一个函数的类型
func add() func(int) int {
	i := 10
	f := func(num int) int { //f等于一个匿名函数
		i += num
		return i
	}
	return f
}

func main() {

	//res := jc(5)
	//fmt.Println(res)

	ffff := add()
	fmt.Println(ffff(2), ffff(2), ffff(2)) //没有影响到i;跳过了初始化i的语句
	dddd1 := add()(2)
	dddd2 := add()(2)
	fmt.Println(dddd1, dddd2) //每次都会调用i:=10,影响i的值;

}

/*
多行注释
*/
