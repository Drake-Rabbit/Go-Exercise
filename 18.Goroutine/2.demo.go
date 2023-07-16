package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var once sync.Once

type singeleton struct {
}

var instance *singeleton

func GetIinstance() *singeleton {
	once.Do(func() {
		fmt.Println("hello once")
		instance = new(singeleton)
	})

	return instance
}

func main() {

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go dec()
	}
	w.Wait()
	fmt.Println(ssum)

	//ins := GetIinstance()
	//ins1 := GetIinstance()
	//fmt.Println(ins1)
	//fmt.Printf("%T", ins)
	//runtime.GOMAXPROCS(1)
	//for i := 0; i <= 10; i++ {
	//	w.Add(1)
	//	go AllCharge(i)
	//}
	//w.Wait()
	//fmt.Println(use)
	//sm.Range(func(key, value any) bool {
	//	fmt.Println(key, value)
	//	return true
	//})

	//sync.Map 为了保证并发安全有一些性能损失

}

var ssum int64 = 1000

func dec() {
	defer w.Done()
	//lock.Lock()
	//ssum -= 1
	//lock.Unlock()
	atomic.AddInt64(&ssum, -1)
}

var w sync.WaitGroup
var lock sync.Mutex //互斥锁
var sm sync.Map

// var lock sync.RWMutexMutex  //读写互斥锁
var sum int
var use = make(map[int]int)

func AllCharge(n int) {
	defer w.Done()
	//lock.Lock()
	sum = 0
	for i := 0; i < n; i++ {
		sum += i
	}
	//use[n] = sum
	//lock.Unlock()
	sm.Store(n, sum)
}
