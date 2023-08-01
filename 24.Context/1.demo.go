package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"net/http"
	"regexp"
)

// form格式
type Pouser struct {
	Name     string `form:"name" binding:"required,mobile,min=4,max=8,alpha"`
	Password string `form:"pass" binding:"required,min=4,max=12"`
}

// json格式
//
//	type Pouser struct {
//		Name     string `json:"name" binding:"required,min=4,max=8,alpha"`
//		Password string `json:"pass" binding:"required,min=4,max=12"`
//	}

type ppuser struct {
	Name     string `form:"name"`
	Password string `form:"pass"`
}
type ppuserrr struct {
	Name     string `uri:"name"`
	Password string `uri:"pass"`
}

func SSum(a, b int) int {
	return a + b
}

func SetCookie(c *gin.Context) {
	c.SetCookie("token", "1321313", 0, "/", "localhost", true, true)
}
func GetCookie(c *gin.Context) {
	token, _ := c.Cookie("token")
	c.String(200, token)
}

// 自定义验证器Validator,检查手机号
func cheeckMobile(f validator.FieldLevel) bool {
	mobile := f.Field().Interface().(string)
	exp := "^1[3|5|6|7|8|9][0-9]\\d{9}$"
	if match, _ := regexp.MatchString(exp, mobile); match {
		return true
	}
	return false
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

	//自定义验证器Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", cheeckMobile)
	}

	//Cookie
	r.GET("setcookie", SetCookie, GetCookie)

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
		//c.HTML(200, "user/index.html", gin.H{
		//	"title": "index.html.......",
		//})

		c.Set("name", "lampo")
		c.String(200, c.GetString("name"))

	})

	r.GET("success", func(c *gin.Context) {
		c.String(200, "login sucess")
	})

	//r.POST("postpouser", func(c *gin.Context) {
	//	var u Pouser
	//	err := c.ShouldBind(&u)
	//
	//	if err != nil {
	//		c.String(200, err.Error())
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"password": u.Password,
	//		"username": u.Name,
	//	})
	//})
	r.POST("looo", func(c *gin.Context) {
		var u Pouser
		err := c.ShouldBind(&u)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"password": u.Password,
			"username": u.Name,
		})
	})

	r.Run()
}
