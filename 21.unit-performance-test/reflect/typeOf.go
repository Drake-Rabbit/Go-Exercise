package main

import (
	"fmt"
	"reflect"
)

type myInt int
type My struct {
	Name string
}

func main() {
	//kind 获得最底层的数据类型
	//Name 表面的数据类型
	user := My{
		Name: "map",
	}
	user2 := &My{
		Name: "map",
	}
	var f myInt = 32
	var d int64 = 20
	aType := reflect.TypeOf(d)
	bType := reflect.TypeOf(f)
	cType := reflect.TypeOf(user)
	dType := reflect.TypeOf(user2)
	fmt.Println("f:", aType, aType.Kind(), aType.Name())
	fmt.Println("d:", bType, bType.Kind(), bType.Name())
	fmt.Println("user:", cType, cType.Kind(), cType.Name())
	fmt.Println("user2:", dType, dType.Kind(), dType.Name()) //Name为空
	//

	//
	var a int32 = 1
	var b float64 = 3.14
	var c string = "hello"
	//fmt.Printf("%T", a)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))
}
