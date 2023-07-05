package main

import "fmt"

func show(arg interface{}) {
	fmt.Printf("\n %T %v \n", arg, arg)
}
func main() {
	//
	var i interface{}
	i = true
	fmt.Printf("%T %v \n", i, i)
	//
	userInfo := make([]interface{}, 2)
	userInfo[0] = "lam"
	userInfo[1] = 45
	fmt.Printf("%T %v \n", userInfo, userInfo)
	//
	info := make(map[interface{}]interface{})
	info[0] = "lamop"
	info["age"] = 45
	fmt.Printf("%T %v \n", info, info)

	//
	show("dadadadad")
}
