package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world"
	in := strings.Index(str, "e")
	inin := strings.LastIndex(str, "o")

	fmt.Println(strings.Contains(str, "wo"))
	fmt.Println(in)
	fmt.Println(inin)
	fmt.Println(strings.Count(str, "l"))
	//大小写
	fmt.Println(strings.ToUpper("w"))
	fmt.Println(strings.ToLower("L"))

}
