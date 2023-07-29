package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func SSum(a, b int) int {
	return a + b
}
func main() {
	r := gin.Default()
	r.Static("img", "24.Context/static/image")
	r.StaticFile("/1.png", "24.Context/static/image/1.png")
	//r.LoadHTMLGlob("24.Context/view/*")
	r.SetFuncMap(template.FuncMap{
		"sum": SSum,
	})
	r.LoadHTMLGlob("24.Context/view/**/*")

	//r.LoadHTMLGlob("24.Context/view/*")
	r.GET("user/index", func(c *gin.Context) {
		//c.JSON(200, gin.H{
		//	"status": "success",
		//	"code":   888,
		//})
		//c.JSONP(200, gin.H{
		//	"status": "success",
		//	"code":   888,
		//})

		//查看GET
		///index?callback=ppp
		//[::1]:12377
		//c.String(200, c.Request.Method+"\n")
		//c.String(200, c.Request.URL.String()+"\n")
		//c.String(200, c.Request.RemoteAddr+"\n")

		//重定向
		//c.Redirect(http.StatusFound, "http://www.baidu.com")
		//
		//c.Request.URL.Path = "/success"
		//r.HandleContext(c)

		//HTML
		c.HTML(200, "user/index.html", gin.H{
			"title": "index.html.......",
		})
	})

	r.GET("success", func(c *gin.Context) {
		c.String(200, "login sucess")
	})

	r.Run()
}
