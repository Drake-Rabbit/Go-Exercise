package main

import (
	"bufio"
	"os"
)

// 我是单行注释
func main() {
	//创建文件夹
	//os.Mkdir("user", os.ModePerm)
	//os.Create("user/user.txt")

	//写入文件
	file, _ := os.OpenFile("user/user.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//file.WriteString("write sucessed\n")
	//file.Write([]byte("os byte\n"))

	//bufio缓冲写入
	write := bufio.NewWriter(file)
	write.WriteString("hello bufio\n") //写入缓冲
	write.Flush()                      //用bufio包  需要flush才能写入文件

	//Ioutil写入  没法追加   覆盖写入
	//os.WriteFile("user/user.txt", []byte("os ioutil\n"), 0666)
}

/*
多行注释
*/
