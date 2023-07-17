package main

import (
	"sync"
)

var p = 0
var wg sync.WaitGroup

func ssu(n int, mut chan bool) {
	defer wg.Done()
	mut <- true
	p += n
	<-mut

}

func main() {
	//单向通道,只能读和只能写
	wch := make(chan<- int, 2)
	rch := make(<-chan string, 2)
	wch <- 5
	x := <-rch

	//ch := make(chan int, 3) //有缓冲通道,可以既写入又输出
	//ch <- 3
	//ch <- 4
	//ch <- 5
	//close(ch)
	//<-ch      //直接扔掉一个值
	//x := <-ch //再扔一个值个给x,只剩一个了
	//fmt.Println(x, len(ch), cap(ch))

	//有缓冲通道使用以及实现互斥锁
	//mut := make(chan bool, 1)
	//for i := 1; i < 2; i++ {
	//	wg.Add(1)
	//	go ssu(i, mut)
	//}
	//fmt.Println(p)

	//两种遍历

	//for {
	//	v, ok := <-ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}

	//for data := range ch {
	//	fmt.Println(data)
	//}
}
