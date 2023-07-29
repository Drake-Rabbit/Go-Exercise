package router

import "github.com/gin-gonic/gin"

func InitttRouter() *gin.Engine {
	router := gin.Default()
	userRouter(router) //r.GET("login", context.String(200, "hello login"))
	//router.GET()
	return router
}
