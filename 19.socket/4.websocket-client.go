package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func getHandle(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("err", err)
	}
	temp.Execute(w, "this is Websocket")
}
func main() {

	http.HandleFunc("/index", getHandle)
	http.ListenAndServe(":7788", nil)
}
