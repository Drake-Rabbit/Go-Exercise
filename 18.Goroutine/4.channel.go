package main

import (
	"fmt"
)

func te(ch chan int) {
	defer wg.Done()
	//x := <-ch
	ch <- 20
	ch <- 21
	ch <- 22
	close(ch)
	//fmt.Println(x)
}

//var wg sync.WaitGroup

// 我是单行注释
func main() {
	//
	var ch1 chan int
	ch1 = make(chan int) //无缓冲通道

	wg.Add(1)
	go te(ch1)
	//x := <-ch1
	//y := <-ch1
	//z := <-ch1
	for data := range ch1 {
		fmt.Println(data)
	}
	x, ok := <-ch1
	fmt.Println(x, ok)
	//ch1 <- 10
	wg.Wait()
	//ch2 <- "string"

	//fmt.Println(ch1, ch2, x)

}

/*
多行注释
*/
