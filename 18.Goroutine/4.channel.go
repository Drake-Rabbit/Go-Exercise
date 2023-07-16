package main

import (
	"fmt"
	"sync"
)

func te(ch chan int) {
	defer wg.Done()
	x := <-ch
	fmt.Println(x)
}

var wg sync.WaitGroup

// 我是单行注释
func main() {
	//
	var ch1 chan int
	ch1 = make(chan int)
	//ch2 := make(chan string)
	//fmt.Println(ch1, ch2)
	wg.Add(1)
	go te(ch1)
	ch1 <- 10
	wg.Wait()
	//ch2 <- "string"

	//fmt.Println(ch1, ch2, x)

}

/*
多行注释
*/
