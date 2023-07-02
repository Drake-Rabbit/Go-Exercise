package main

import "fmt"

// 我是单行注释

type User struct {
	name           string
	age            int
	email, address string
}

func main() {

	//匿名结构体
	var p struct {
		name string
		age  int
	}
	p.name = "匿名"
	p.age = 22
	fmt.Println(p, "-----")
	fmt.Printf("\n%T\n", p)

	//user := new(User)
	//var user User
	//user := &User{}
	//user.name = "lampol"
	//user.age = 23
	//user.email = "12312@aa.com"
	//user.address = "lasdhkashd"

	user1 := &User{
		name:    "lampol",
		age:     23,
		email:   "12312@aa.com",
		address: "lasdhkashd",
	}

	user2 := User{
		name:    "lampol",
		age:     23,
		email:   "12312@aa.com",
		address: "lasdhkashd",
	}

	fmt.Println(user1, "-----")
	fmt.Printf("\n%T\n", user1)
	fmt.Println(user2, "-----")
	fmt.Printf("\n%T\n", user2)
}

/*
多行注释
*/
