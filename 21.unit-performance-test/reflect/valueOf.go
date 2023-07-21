package main

import (
	"fmt"
	"reflect"
)

func main() {

	//反射修改它的值
	var c int = 66
	cValue := reflect.ValueOf(&c)
	cValue.Elem().SetInt(999)
	//a = 100
	fmt.Println(c)
	//
	var a int = 88
	var b int
	aValue := reflect.ValueOf(a)
	fmt.Println(aValue)
	fmt.Printf("%T\n", aValue)
	fmt.Printf("%T\n", aValue.Int()) //aValue.Int() 返回int形式
	b = aValue.Interface().(int)
	fmt.Printf("%T\n", aValue.Interface())
	fmt.Println(b)
}
