package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 我是单行注释
func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe("127.0.0.1:7788", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	//temp, err := template.New("test").Parse("<h1> {{.}} </h>")
	temp, err := template.ParseFiles("17.net-http---html-tamplate/index.tpl.html")
	if err != nil {
		fmt.Println("err:", err)
	}
	user := make(map[string]string)
	user["name"] = " lam"
	user["email"] = "35435@qq.com"
	user["age"] = "50"

	temp.Execute(w, user)
}

/*
多行注释
*/
