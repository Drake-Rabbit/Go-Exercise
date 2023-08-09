package main

import "Go-Exercise/25.verification-code/gin-for-captview/router"

func main() {
	router := router.InitRouter()
	router.Run(":8777")
}
