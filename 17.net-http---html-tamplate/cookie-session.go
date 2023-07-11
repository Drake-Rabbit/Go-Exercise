package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"strconv"
)

// 我是单行注释
func main() {
	http.HandleFunc("/setSession", setSession)
	http.HandleFunc("/getSession", getSession)

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.ListenAndServe("127.0.0.1:7788", nil)
}

var store = sessions.NewCookieStore([]byte(os.Getenv("4651646516")))

func getSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "LSDHFJKLSDHFLDS")
	age := session.Values["id"]
	name := session.Values["user_name"]
	w.Write([]byte(name.(string)))
	w.Write([]byte(strconv.Itoa(age.(int))))
}

func setSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "LSDHFJKLSDHFLDS")
	session.Values["user_name"] = "fanbingb"
	session.Values["id"] = 5
	session.Save(r, w)
	w.Write([]byte("set Session success"))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := &http.Cookie{
		Name:  "name",
		Value: "guofucheng",
	}

	cookie2 := &http.Cookie{
		Name:  "age",
		Value: "66",
	}

	http.SetCookie(w, cookie1)
	http.SetCookie(w, cookie2)
	w.Write([]byte("save cookie success"))
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Cookie("name")
	age, _ := r.Cookie("age")

	for _, c := range r.Cookies() {
		fmt.Println(c.Value)
	}

	w.Write([]byte(name.Value))
	w.Write([]byte(age.Value))
}

/*
多行注释
*/
