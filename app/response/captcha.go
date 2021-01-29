package response

type ImgCaptcha struct {
	Response
	Data Data `json:"data"`
}

type Data struct {
	Response
	CaptchaID string `json:"captchaId"`
	LocX      int    `json:"locX"`
	LocY      int    `json:"locY"`
}