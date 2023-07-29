package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// 获取参数
func main() {

	r := gin.Default()
	//必选参数 /user/:ID  只能匹配/user/60
	//使用 :param 的语法完成必选参数，例如 /user/:ID，即可匹配 /user/60 这个 URI，但不能匹配 /user

	//r.GET("user/:id", func(c *gin.Context) {
	//	fmt.Println(c.Param("id"))   //后台输出
	//	c.String(200, c.Param("id")) //浏览器输出
	//})

	//可选参数  user/*id 可以匹配/user/60 和/user
	r.GET("user/*id", func(c *gin.Context) {
		fmt.Println(strings.Trim(c.Param("id"), "/"))   //后台输出
		c.String(200, strings.Trim(c.Param("id"), "/")) //浏览器输出
		//Trim 去掉"/"
	})

	//运行
	//r.Run()
}

/*
多行注释
*/
