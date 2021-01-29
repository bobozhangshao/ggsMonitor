package app

import (
	"bytes"
	"fmt"
	"ggsMonitor/utils/app"
	"ggsMonitor/utils/request"
)

// @Description   发送手机验证码
// @Auth          Qap  16:15 2021/01/26
// @Param         输入参数名        string         手机号
// @Return        返回参数名        string         验证码id
func SendToMobile(mobile, captchaId string) {
	jsonStr := app.H{"mobile": mobile, "type": "LOGIN_CONFIRMATION", "captchaId": captchaId}
	data := bytes.NewBuffer([]byte(app.MapToJson(jsonStr)))
	res, _ := request.Post("https://prod.ggszhg.com/xgt-app/applet/personalCenter/sendToMobile?os=APPLET&osVersion=1.0.0&userId=&userToken=&sign=D608393BFA0A548C0B77DE5927E52170", data)
	fmt.Println(res.String())
}

// @Description   登录
// @Auth          Qap  16:14 2021/01/26
// @Param         输入参数名        string         手机号
// @Return        返回参数名        string         验证码
func Login(phone, verifyCode string) {
	jsonStr := app.H{"accountTypes": "MEMBER", "applicationType": "SMALL_PROGRAM", "phone": phone, "verifyCode": verifyCode, "memberType": "COMMON", "verifyType": "VERIFY_CODE", "authCode": "093jHLkl2oPJo64qgjll2wStA40jHLkg"}
	data := bytes.NewBuffer([]byte(app.MapToJson(jsonStr)))
	res, _ := request.Post("https://prod.ggszhg.com/xgt-app/applet/personalCenter/login?os=APPLET&osVersion=1.0.0&userId=&userToken=&sign=EF113FD3D066F82060BB46C1C4A1AAAD", data)
	fmt.Println(res.String())
}