package main

import (
	"github.com/gin-gonic/gin"
)

func loginMiddleWare(c *gin.Context) {
	c.String(200, "loginmiddleware....\n")
	c.Next() //先跳过中间件，执行完母函数的handle，然后再回来继续执行
	c.String(200, "login is cpmpleted....\n")

}

func logoutMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "loginOUT....\n")
	}
}
func main() {
	//gin.Default() 初始化路由对象时，会随之附加两个中间件 Logger 和 Recovery
	//gin.New()  创建空白路由  没有任何中间

	//r := gin.Default()
	r := gin.New()
	//r.Use(loginMiddleWare, logoutMiddleWare()) //中间件往这里放就完了
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "hello")
	//})
	usepp := r.Group("userpp")
	{
		usepp.GET("lgin", loginMiddleWare, func(c *gin.Context) {
			c.String(200, "login usepp\n")
		})

	}
	r.Run()
}
