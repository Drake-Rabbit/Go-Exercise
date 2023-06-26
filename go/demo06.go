package main

func main() {
	//a =100
	// b =200
	//temp =0
	a := 100
	//var a int = 100
	var b int = 99

	a, b = b, a
	println(a, b)
}
