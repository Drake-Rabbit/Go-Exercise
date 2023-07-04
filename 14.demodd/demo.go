package main

import "fmt"

const timeContext = "2006-01-02 08:02:12"

// 我是单行注释
func main() {
	//pay.AliPay()
	//fmt.Println(pay.AppId)
	//fmt.Println(time.Now().Format(timeContext))
	//fmt.Println(time.Now())
	//t := time.Now()
	//fmt.Println(time.Now().Unix())
	//fmt.Println(t.Year(), t.YearDay(), t.Day(), time.Hour, t.Minute())
	//
	//st := 163598485
	//fmt.Println(time.Unix(int64(st), 0).Format(timeContext))
	//
	//fmt.Println(time.Now().Add(time.Hour).Format(timeContext))

	//tick := time.NewTicker(2 * time.Second) //计时器
	//for _ = range tick.C {
	//	fmt.Println("hello ")
	//}

	//n := 5
	//for {
	//	time.Sleep(2 * time.Second)
	//}

	timetest()
	fmt.Println("hp")
}
func timetest() {

	defer func() { //必须配合defer
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic---------")
	//var str []string
	//str[0] = "hello"
}

/*
多行注释
*/
