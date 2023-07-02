package main

import "fmt"

// 我是单行注释
type User4 struct {
	name    string
	lan     []string
	address map[string]string
}

func main() {
	var u User4
	u.name = "lam"
	u.lan = make([]string, 2, 5)
	u.lan[0] = "php"
	u.lan[1] = "go"
	u.address = make(map[string]string)
	u.address["pro"] = "sss"
	u.address["orp"] = "ddd"
	fmt.Println(u)
	fmt.Printf("%#v,\n%T", u, u)

}

/*
多行注释
*/
