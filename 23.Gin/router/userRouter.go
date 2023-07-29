package router

import (
	"Go-Exercise/23.Gin/controller"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.Engine) {
	r.GET("login", controller.Login)
}
