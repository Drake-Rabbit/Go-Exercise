package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleC(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err", err)
	}
	go handleMsg(conn)

}
func handleMsg(conn *websocket.Conn) {
	defer conn.Close()
	for {

		mType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("err", err)
		}
		//fmt.Println(mType, string(msg))
		conn.WriteMessage(mType, msg)
	}
}
func main() {

	http.HandleFunc("/websocket", handleC)
	http.ListenAndServe(":5566", nil)

}
