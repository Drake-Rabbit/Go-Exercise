package main

import "fmt"

// 我是单行注释
type Userq struct {
	name           string
	age            int
	email, address string
}

func main() {
	var u4 = Userq{
		"dasda",
		23,
		"113@aa.com",
		"dada",
	}
	fmt.Printf("%#v,%T", u4, u4)
	fmt.Println()

	u1 := Userq{name: "lam", age: 35}
	fmt.Printf("%#v,%T", u1, u1)

	//加地址符号，返回指针
	fmt.Println()
	u2 := &Userq{name: "lam", age: 35}
	fmt.Printf("%#v,%T", u2, u2)

	fmt.Println()
	//new 实例化
	u3 := new(Userq)
	u3.name = "u3u3"
	u3.age = 3
	fmt.Printf("%#v,%T", u3, u3)

}

/*
多行注释
*/
