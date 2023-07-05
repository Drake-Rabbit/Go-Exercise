package main

import "fmt"

func main() {
	//类型断言（Type Assertion）是一个使用在接口值上的操作，
	//用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。
	//value, ok := x.(T)
	//
	//类型断言就是将接口类型的值(x)，转换成类型(T)。
	//
	//返回两个参数，第一个参数是x转化为T类型后的变量，
	//第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败

	var i interface{}
	i = 23
	value, ok := i.(int)
	fmt.Println("value:", value, ok)
	//
	v2 := i.(int)
	fmt.Println(v2)
	//

	var h interface{}
	h = 23
	//if _, ok := h.(int); ok {
	//	fmt.Println("h is int")
	//} else if _, ok := i.(string); ok {
	//	fmt.Println("h is string")
	//}

	//查询一个接口变量绑定的底层变量类型
	switch h.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("string")
	}

}
