package main

//
//import (
//	"fmt"
//)
//
//// 接口的声明
//type Mobiler interface {
//	Call(i int)
//	Send(string2 string) bool
//}
//
//type Mobiler2 interface {
//	Call(i int)
//}
//
//type Huawei struct {
//	Name string
//}
//
//func (h Huawei) Call(mo int) {
//	fmt.Println(h.Name, " call phone", mo)
//}
//
//func (h Huawei) Send(mo string) bool {
//	fmt.Println(mo)
//	return true
//}
//func (h Huawei) playGame() {
//	fmt.Println("play game")
//}
//
//type Xiaomi struct {
//	Name string
//}
//
//func (x Xiaomi) Call(mo int) {
//	fmt.Println(x.Name, "call phone", mo)
//}
//
//func main() {
//	var m Mobiler
//	//var h = &Huawei{"p40"}
//	m = &Huawei{"p40"} //1.用接口方式调用----对象实例赋值给接口
//	//h.playGame()
//	//h.Call(110) 实例化调用
//
//	//1.用接口方式调用----对象实例赋值给接口
//	//m = h
//	m.Call(122222222)
//	m.Send("546")
//	//2.把接口赋值给另一个接口
//	var m2 Mobiler = &Huawei{"p60"}
//	m2 = m
//	m2.Send("Two-接口")
//
//	var m3 Mobiler2 = Xiaomi{"K60"}
//	m3 = m2 //只能方法多的接口才能赋值给另一个方法少的接口
//	m3.Call(555)
//}
//
///*
//多行注释
//*/
