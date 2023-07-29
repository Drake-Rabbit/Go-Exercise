package controller

import "github.com/gin-gonic/gin"

func Login(context *gin.Context) {
	context.String(200, "hello controller login")
}
