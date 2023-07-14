package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {
	//runtime.GOMAXPROCS(1)
	for i := 0; i <= 10; i++ {
		w.Add(1)
		go AllCharge(i)
	}
	w.Wait()
	fmt.Println(use)
}

var lock sync.Mutex //互斥锁

// var lock sync.RWMutexMutex  //读写互斥锁
var sum int
var use = make(map[int]int)

func AllCharge(n int) {
	defer w.Done()
	lock.Lock()
	sum = 0
	for i := 0; i < n; i++ {
		sum += i
	}
	use[n] = sum
	lock.Unlock()
}
