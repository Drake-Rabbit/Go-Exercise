package main

import "fmt"

// 我是单行注释

type a struct {
	name    string
	address address
}
type address struct {
	mobile string
	pro    string
	city   string
}

func main() {
	var u a
	u.name = "lam"
	u.address.mobile = "138888"
	u.address.pro = "138888"
	u.address.city = "138888"
	fmt.Printf("%#v,\n%T", u, u)

}

/*
/*
多行注释
*/
