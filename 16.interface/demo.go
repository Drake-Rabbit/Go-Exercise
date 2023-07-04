package main

import "fmt"

// 接口的声明
type Mobiler interface {
	Call(i int)
}
type Huawei struct {
	Name string
}

func (h Huawei) Call(mo int) {
	fmt.Println(h.Name, " call phone", mo)
}

func main() {
	var m Mobiler
	var h = Huawei{
		Name: "p40",
	}
	//h.Call(110) 实例化调用
	//用接口调用
	m = h
	m.Call(12222)

}

/*
多行注释
*/
