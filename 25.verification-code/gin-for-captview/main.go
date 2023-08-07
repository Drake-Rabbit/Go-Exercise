package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"net/http"
)

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("capt", verifyCaptcha)
	}
	router.LoadHTMLGlob("captview/*")
	router.GET("user", index)
	router.GET("/api/getCaptcha", getCaptcha)
	router.POST("/login", Login)
	router.Run(":8777")
}

func verifyCaptcha(f validator.FieldLevel) bool {
	captcha := f.Field().FieldByName("Captcha").Interface().(string)
	captcha_id := f.Field().FieldByName("Captcha_id").Interface().(string)

	if store.Verify(captcha_id, captcha, true) {
		return true
	}
	return false
}

type causer struct {
	Captcha   string `form:"captcha" binding:"required,capt"`
	CaptchaId string `form:"captcha_id" binding:"required"`
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

var store = base64Captcha.DefaultMemStore

func getCaptcha(context *gin.Context) {
	driverString := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},

		Fonts: []string{"wqy-microhei.ttc"},
	}
	var driver base64Captcha.Driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//json.NewEncoder(w).Encode(body)
	context.JSON(200, body)
}

func index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
