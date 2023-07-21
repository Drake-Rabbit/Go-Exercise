package main

import (
	"fmt"
	"reflect"
)

type Uyser struct {
	Name string `json:"name"`
	Age  int
}

func (u Uyser) run() {
	fmt.Println(u.Name, "user is run ")
}

func (u Uyser) sleep(add string) {
	fmt.Println(u.Name, "user is sleep ")
}

// 我是单行注释
func main() {

	uyser := Uyser{"mapdad", 23}
	uyType := reflect.TypeOf(uyser)
	uyValue := reflect.ValueOf(uyser)
	//
	fmt.Println(uyType.NumMethod(), uyValue.NumMethod(), "\n")
	//
	for i := 0; i < uyValue.NumMethod(); i++ {
		fmt.Println(uyType.Method(i).Name)
		var arg = []reflect.Value{}
		arg = append(arg, reflect.ValueOf("南京"))
		uyValue.Method(i).Call(arg)
	}

	//
	//fmt.Println(uType.Kind(), uType.Name(), uType.NumField())
	//fmt.Println(uType.Field(0))
	//fmt.Println(uType.Field(1))
	//fmt.Println()
	//for i := 0; i < uType.NumField(); i++ {
	//	field := uType.Field(i)
	//	fmt.Println(field.Tag, field.Name, field.Index)
	//}
	//fmt.Println()
	////根据给定字符串返回字符串对应的结构体字段的信息
	//ageFile, _ := uType.FieldByName("Age")
	//fmt.Println("ageFile:", ageFile)
	//fmt.Println("ageFile.Name:", ageFile.Name)
}
