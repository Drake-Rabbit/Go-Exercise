package main

import "fmt"

type U7 struct {
	name string
}

// 需要传指针类型的接受者，才会去同步改变实例化对象的值
func (u *U7) setname(name string) {
	u.name = name
}

// 自定义类型的
type mystring string

func (m mystring) echo(str string) {
	fmt.Println(str)
}
func main() {

	var u U7
	u.name = "lam"
	u.setname("loool")
	fmt.Println(u, "\n")

	var m mystring
	m.echo("hello mystring")

}

/*
多行注释
*/
