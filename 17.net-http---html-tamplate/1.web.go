package main

import (
	"fmt"
	"net/http"
)

func getHandle(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("hello http")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL)
	fmt.Println("body:", r.Body)
	fmt.Println("header:", r.Header)
	fmt.Println("RemoteAddr:", r.RemoteAddr)
	//w.Write([]byte("hello http"))
}

type myHandler struct {
}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello serverHttp"))
}

// 我是单行注释
func main() { //对127.0.0.1:6666/index进行监听
	//http.HandleFunc("/index", getHandle)
	//http.ListenAndServe("127.0.0.1:7788", nil)

	var handler myHandler
	//自定义server
	server := http.Server{
		Addr:    ":7788",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("err", err)
	}
}

/*
多行注释
*/
