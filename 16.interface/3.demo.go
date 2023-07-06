package main

import "fmt"

// 接口的嵌套
type pc interface {
	game
	internet
}

type game interface {
	playgame()
}
type internet interface {
	callint(int)
}
type huaw struct {
	name string
}

func (h huaw) callint(mo int) {
	fmt.Println(h.name, "get internet", mo)
}
func (h huaw) playgame() {
	fmt.Println(h.name, "is playing game")
}
func main() {
	var m pc = huaw{"P40"}
	m.callint(564654654)
	m.playgame()

}

/*
多行注释
*/
