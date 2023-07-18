package main

import (
	"fmt"
	"time"
)

// 超时处理
func testt(ch3 chan int) {
	time.Sleep(100 * time.Second) //等待100秒
	ch3 <- 12
	close(ch3)
}
func main() {
	ch3 := make(chan int)
	go testt(ch3)
	select {
	case data := <-ch3:
		fmt.Println(data)
	// 超时处理
	case <-time.After(time.Second): //等待了一秒
		fmt.Println("我等待了一秒了，不等你了")
	}

	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	ch1 <- 20
	ch1 <- 30
	ch2 <- "hello"
	ch2 <- "cao"

	for {

		select {
		case a := <-ch1:
			fmt.Println(a)
		case b := <-ch2:
			fmt.Println(b)

		default:
			return
		}
	}
}
