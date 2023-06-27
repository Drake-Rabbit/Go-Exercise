package main

import "fmt"

func main() {
	//var num1 []int
	var num1 = make([]int, 5) //[5 6 7 0 0]
	//num1:=[]int{1,2,3,4} //[5 6 7 4]
	//num1:=[]int{1,2,3} //[5 6 7]
	// num1:=[]int{1,2} //[5 6]
	num2 := []int{5, 6, 7}
	count := copy(num1, num2)
	fmt.Println(num1, num2, count)

}
