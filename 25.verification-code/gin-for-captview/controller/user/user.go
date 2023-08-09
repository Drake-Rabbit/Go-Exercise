package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type causer struct {
	Captcha   string `form:"captcha" binding:"required,capt"`
	CaptchaId string `form:"captcha_id" binding:"required"`
}

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func Login(context *gin.Context) {
	var u causer
	err := context.ShouldBind(&u)
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
		context.String(200, "code ok")
	}

}
