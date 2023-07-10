package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHandle(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("hello http")
	//fmt.Println("method:", r.Method)
	//fmt.Println("url:", r.URL)
	//fmt.Println("body:", r.Body)
	//fmt.Println("header:", r.Header)
	//fmt.Println("RemoteAddr:", r.RemoteAddr)
	w.Write([]byte("hello http"))
}

type myHandler struct {
}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello serverHttp"))
}
func get(w http.ResponseWriter, request *http.Request) {
	u := "http://127.0.0.1:7788/index"
	resp, err := http.Get(u)

	fmt.Println(resp.Body)
	body, _ := io.ReadAll(resp.Body) //ioutil 已经弃用
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(body))

}
func getparams(w http.ResponseWriter, r *http.Request) {
	//defer r.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("id"), data.Get("name"))
	w.Write([]byte("hello getparams"))
}

func params(w http.ResponseWriter, r *http.Request) {
	u := "heep://127.0.0.1:7788/getParams?name=lampol&id=5"
	//defer resp.Close()
	resp, err := http.Get(u)
	fmt.Println(resp.Body)
	body, _ := io.ReadAll(resp.Body) //ioutil 已经弃用
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(string(body))
}
func post(w http.ResponseWriter, r *http.Request) {
	//u := "http://127.0.0.1:7788/getpost"
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=liming&age=23"
	//resp, err := http.Post(u, contentType, strings.NewReader(data))

	u := "http://127.0.0.1:7788/getpost"
	contentType := "application/json"
	data := `{"name":"zzz","age":"66"}`
	resp, err := http.Post(u, contentType, strings.NewReader(data))

	if err != nil {
		fmt.Println("err", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body) //ioutil 已经弃用
	w.Write([]byte(string(body)))

}

type uiui struct {
	Name any
	Age  any
}

func getpost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body) //ioutil 已经弃用
	if err != nil {
		fmt.Println("err", err)
	}
	var uiui uiui
	json.Unmarshal(body, &uiui)
	fmt.Println(string(body))
	fmt.Println(uiui.Age, uiui.Name)
	//r.ParseForm()
	//name := r.PostFormValue("name")
	//age := r.PostFormValue("age")
	//fmt.Println(name, age)
	w.Write([]byte("getpost"))
}

// 我是单行注释
func main() { //对127.0.0.1:6666/index进行监听
	//http.HandleFunc("/index", getHandle)
	//http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
	http.HandleFunc("/getpost", getpost)

	http.HandleFunc("/params", params)
	http.HandleFunc("/getparams", getparams)
	http.ListenAndServe("127.0.0.1:7788", nil)

	//var handler myHandler
	////自定义server
	//server := http.Server{
	//	Addr:    ":7788",
	//	Handler: handler,
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	fmt.Println("err", err)
	//}
}

/*
多行注释
*/
