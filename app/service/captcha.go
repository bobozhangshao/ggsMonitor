package service

import (
	"bytes"
	"fmt"
	"ggsMonitor/app/response"
	"ggsMonitor/utils/app"
	"ggsMonitor/utils/request"
)

// @获取滑块验证
// @Author Qap
// @Date 14:01 2021/01/26
func GetImgCaptcha() {
	var resp response.ImgCaptcha
	message := bytes.NewBuffer([]byte(""))
	res, _ := request.Post("https://prod.ggszhg.com/xgt-app/applet/imgCaptcha/getImgCaptcha?os=APPLET&osVersion=1.0.0&userId=&userToken=&sign=9F00FD42321B5FB62FAA30E72D59DDB9", message)
	res.BindJson(&resp)
	fmt.Println(resp)
	ValidateCaptcha(resp.Data.CaptchaID, resp.Data.LocX)
}

// @验证滑块
// @Author Qap
// @Date 14:01 2021/01/26
func ValidateCaptcha(captchaId string, locX int) {
	requestData := app.H{"captchaId": captchaId, "LocX": locX + 1}
	data := app.MapToJson(requestData)
	message := bytes.NewBuffer([]byte(data))
	res, _ := request.Post("https://prod.ggszhg.com/xgt-app/applet/imgCaptcha/validateCaptcha?os=APPLET&osVersion=1.0.0&userId=&userToken=&sign=557B4BC71A5B5CA5E8B2FE1FD1A7AD7E", message)
	fmt.Println(res.String())
}