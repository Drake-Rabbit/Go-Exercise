package main

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"html/template"
	"image/color"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/index", index)
	//
	//
	//api for create captcha
	http.HandleFunc("/api/getCaptcha", generateCaptchaHandler)

	//api for verify captcha
	http.HandleFunc("/api/verifyCaptcha", verifyCaptchaHandler)

	//fmt.Println("Server is at :8777")
	//if err := http.ListenAndServe(":8777", nil); err != nil {
	//	log.Fatal(err)
	//}
	http.ListenAndServe(":8777", nil)
}

func verifyCaptchaHandler(writer http.ResponseWriter, request *http.Request) {
	captcha := request.PostFormValue("captcha")
	captcha_id := request.PostFormValue("captcha_id")

	body := map[string]interface{}{"code": 0, "msg": "failed"}
	if store.Verify(captcha_id, captcha, true) {
		body = map[string]interface{}{"code": 1, "msg": "ok"}
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(writer).Encode(body)

	//fmt.Println(captcha, captcha_id)
}

func index(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("captview/index.html")
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.Execute(writer, nil)
}

func captchaVerifyHandle(writer http.ResponseWriter, request *http.Request) {

}

var store = base64Captcha.DefaultMemStore

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
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
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(body)
}
