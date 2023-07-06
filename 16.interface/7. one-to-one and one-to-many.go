package main

import "fmt"

//  双接口被一结构体实现

type I interface {
	send(string2 string)
}

type Ga interface {
	playgame()
}

type Huaa struct {
	Name string
}

func (h Huaa) send(str string) {
	fmt.Println(h.Name)
}

func (h Huaa) playgame() {
	fmt.Println("game")
}

// 但接口只能全部被实现，但必须有一个结构体完全实现此接口的所有方法才行，否则会报错
func main() {

	var ii I = Huaa{"ppp30"}
	ii.send("dad")
	var g Ga = Huaa{"ppp"}
	g.playgame()
}

/*
多行注释
*/
