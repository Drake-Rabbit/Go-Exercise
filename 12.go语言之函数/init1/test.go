package init1

import "fmt"

func init() {
	fmt.Println("init1-----test我是先引入的init包")
}
func Add() {
	fmt.Println("adddd")
}
