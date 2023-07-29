package main

import (
	"Go-Exercise/23.Gin/router"
	"github.com/gin-gonic/gin"
)

// 路由分组
func main() {
	r := router.InitttRouter()
	r.Run()
	//r := gin.Default()
	//user := r.Group("user")
	//{
	//	user.GET("login", login2)
	//	user.GET("logout", logout2)
	//}
	//
	//goods := r.Group("goods")
	//{
	//	goods.GET("list", list2)
	//	goods.GET("add", add2)
	//}
	//r.Run()
}

func add2(context *gin.Context) {
	context.String(200, "goods add2")
}

func list2(context *gin.Context) {
	context.String(200, "goods list2")
}

func logout2(context *gin.Context) {
	context.String(200, "user logout2")
}

func login2(context *gin.Context) {
	context.String(200, "user login2")
}
