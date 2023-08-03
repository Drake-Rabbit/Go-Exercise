package main

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"html/template"
	"image/color"
	"net/http"
)

func main() {

	http.HandleFunc("/index", index)
	//

	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", captchaVerifyHandle)

	//fmt.Println("Server is at :8777")
	//if err := http.ListenAndServe(":8777", nil); err != nil {
	//	log.Fatal(err)
	//}
	http.ListenAndServe(":8777", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("captview/index.html")
	t.Execute(writer, nil)
}

func captchaVerifyHandle(writer http.ResponseWriter, request *http.Request) {

}

var store = base64Captcha.DefaultMemStore

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {

	drviceString := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "123456789ABCDEFGHIJKMNOPQRSTUVWXYZ",
		BgColor: &color.RGBA{
			R: 5,
			G: 10,
			B: 23,
			A: 10,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	var driver base64Captcha.Driver = drviceString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
