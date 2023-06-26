package main

import (
	"fmt"
	"strings"
)

func main() {

	//str := "         dhello  world    "
	//strr := "dhello,world"

	//fmt.Println(strings.Trim(str, "d"))
	//fmt.Println(strings.TrimRight(str, "d"))
	//fmt.Println(strings.TrimSpace(str))

	//res := strings.Split(strr, ",")
	//fmt.Printf("%T", res)
	//fmt.Println()
	//fmt.Println(res)
	//fmt.Println(res[0])

	ss := "hh"
	fmt.Println(strings.Repeat(ss, 5))
	fmt.Println(strings.Replace(ss, "h", "a", 1))
	a := "aaaa"
	A := "AAAA"
	fmt.Println(strings.EqualFold(a, A)) //忽略大小写比较
}
