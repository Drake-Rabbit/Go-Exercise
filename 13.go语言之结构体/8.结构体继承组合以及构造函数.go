package main

import "fmt"

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

func main() {
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
