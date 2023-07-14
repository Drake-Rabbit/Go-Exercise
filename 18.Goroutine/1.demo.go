package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// 我是单行注释
func test1() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("test1()", i)
	}
}

func test2() {
	defer wg.Done()
	for i := 1; i < 5; i++ {
		fmt.Println("test2()", i)
	}
}

func main() {

	//runtime.GOMAXPROCS(1)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	wg.Add(2)
	go test1()
	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	go test2()
	wg.Wait()
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("main()...")

	//go func() {
	//	fmt.Println("hello gproutine")
	//}()

	//time.Sleep(time.Second)

}

/*
多行注释
*/
