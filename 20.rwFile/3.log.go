package main

import (
	"log"
	"os"
)

func main() {

	//log.Fatalln("hello fatal")
	// 输出到文件
	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	log.SetOutput(file)
	log.SetPrefix("[USER]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	log.Println("hello print")

}
