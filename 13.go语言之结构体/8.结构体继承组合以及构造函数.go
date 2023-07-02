package main

import "fmt"

type Usere struct {
	name string
	age  int
}
type Person struct {
	name string
	age  int
}

type male struct {
	height int
	per    Person
}

// male 继承 person的方法
func (p Person) run(h int) {
	fmt.Println(p.name, p.age, h, "runnnn")
}

// 实现构造函数 Test a = new Test();//调用构造函数
func newUser(name string, age int) *Usere {
	return &Usere{
		name: name,
		age:  age,
	}
}

func main() {
	user := newUser("lammmm", 33)
	fmt.Println(user.name)

	m := male{
		height: 170,
		per: Person{
			name: "lam",
			age:  23,
		},
	}
	m.per.run(m.height)
	fmt.Println()
}

/*
多行注释
*/
