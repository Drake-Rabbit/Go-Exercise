package main

import "fmt"

/**
* @Description:
* @Author: oDuck
* @Date: 2023/6/27
 */
func main() {
	//ptr := new(type)
	//使用new方法来定义指针
	ptr := new(string)
	*ptr = "hello"
	ab := *ptr
	fmt.Println(ab)

	a, b := 100, 200
	t := *(&a)
	*&a = *(&b)
	*&b = t

	fmt.Println("a:", a, "b:", b, *&a)

	//num := 100
	//ptr := &num
	//*ptr = 200
	//fmt.Println(num)

}
