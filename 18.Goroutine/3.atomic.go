package main

import (
	"fmt"
	"sync/atomic"
)

var sss int64 = 0
var a, b, c int64 = 22, 22, 44

func main() {
	//如果a与b相等，则c的值给a
	atomic.CompareAndSwapInt64(&a, b, c)
	fmt.Println(a, b, c)
	//直接交换，b的值给a
	atomic.SwapInt64(&a, b)
	fmt.Println(a, b, c)

	var i int64
	for i = 0; i < 10; i++ {
		wg.Add(2)
		go wsum(i)
		go rsum()
	}
	wg.Wait()
}

//var wg sync.WaitGroup

func rsum() {
	defer wg.Done()
	fmt.Println("sss", atomic.LoadInt64(&sss))
}

func wsum(n int64) {
	defer wg.Done()
	atomic.StoreInt64(&sss, n)
}
