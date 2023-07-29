package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理器
func index(c *gin.Context) {
	c.String(http.StatusOK, "hello index")
}
func login(context *gin.Context) {
	context.String(200, "hello login")
}
func update(context *gin.Context) {
	context.String(200, "hello update")
}
func del(context *gin.Context) {
	context.String(200, "hello del")
}
func any(context *gin.Context) {
	context.String(200, "hello any")
}

// 我是单行注释
func main() {
	//处理器分为两类，中间件和请求处理器（业务逻辑处理器）。
	//在 router.GET() 类的的方法中，
	//最后一个 handler 就是请求处理器，除此之外前边的都是中间件，必须要有一个 handler 才可以
	r := gin.Default()
	r.GET("index", index)
	r.POST("login", login)
	r.PUT("update", update)
	r.DELETE("del", del)
	r.Any("any", any) //所有的any请求post，put，post都可以进来
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "没有找到路由")

	})
	//r.Run()

	//r.GET("/", func(c *gin.Context) {
	//
	//	//返回JSON型数据
	//	//c.JSON(200, gin.H{
	//	//	"name": "lampo",
	//	//	"age":  18,
	//	//})
	//	//返回String类型数据
	//	c.String(200, "Hello Gin")
	//})

}

/*
多行注释
*/
