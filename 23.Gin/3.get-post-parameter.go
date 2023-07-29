package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//获取Get参数
	//c.Query(“name”)    获取name值
	//c.DefaultQuery(“age”, “55”)  获取age值没有默认55
	//c.Request.URL.Query()   获取get所有参数
	r.GET("user", func(c *gin.Context) {
		//c.String(200, c.Query("id"))
		//c.String(200, c.Query("name"))
		//c.String(200, c.DefaultQuery("age", "age=404"))
		data := c.Request.URL.Query()
		c.JSON(200, data)
	})
	//c.PostForm(“name”)  获取name
	//c.Request.PostForm  获取所有post参数
	//c.DefaultPostForm(“age”, 55)  默认 55

	r.POST("login", func(c *gin.Context) {
		//name := c.PostForm("name")
		//pass := c.PostForm("password")
		//pass := c.PostForm("pass")

		//c.String(http.StatusOK, name)
		//c.String(200, pass)
		c.String(200, c.DefaultPostForm("age", "99"))

		data := c.Request.PostForm
		c.JSON(200, data)
	})
	//r.Run()
}
