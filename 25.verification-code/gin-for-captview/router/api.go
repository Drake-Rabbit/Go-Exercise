package router

import (
	"Go-Exercise/25.verification-code/gin-for-captview/common/captcha"
	"Go-Exercise/25.verification-code/gin-for-captview/controller/user"
	"Go-Exercise/25.verification-code/gin-for-captview/upload"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("capt", captcha.VerifyCaptcha)
	}
	router.Static("upload", "uplaod")
	router.LoadHTMLGlob("captview/*")
}
func InitRouter() *gin.Engine {
	router.GET("user", user.Index)
	router.GET("/api/getCaptcha", captcha.GetCaptcha)
	router.POST("/login", user.Login)
	router.POST("upload", upload.Upload)
	return router
}
