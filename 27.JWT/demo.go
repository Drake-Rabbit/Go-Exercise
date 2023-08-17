package main

import (
	"fmt"
	//"Go-Exercise/25.verification-code/gin-for-captview/router"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("getToken", getToken)
	router.POST("parseToken ", parseToken)
	router.Run()

}

type User struct {
	Token string `form:"token" binding:"required"`
}

func parseToken(c *gin.Context) {
	var u User
	err := c.ShouldBind(&u)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	fmt.Println(u.Token)
	clanims := UserClaims{}
	token, err := jwt.ParseWithClaims(u.Token, &clanims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKEY), nil
	})
	c.JSON(200, gin.H{
		"token": token,
	})
}

type UserClaims struct {
	UserName string
	jwt.StandardClaims
}

var JWTKEY = "123456789"

func getToken(c *gin.Context) {
	claims := UserClaims{
		UserName: "lampol",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Second).Unix(), //过期时间：当前时间+60s
			IssuedAt:  time.Now().Unix(),                       //发布时间
			Subject:   "user token",
			Issuer:    "127.0.0.1",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, err := token.SignedString([]byte(JWTKEY))
	if err != nil {
		log.Fatal(err.Error())
	}
	c.JSON(200, gin.H{
		"code":  400,
		"msg":   "token sucess",
		"token": tokenString,
	})
}
