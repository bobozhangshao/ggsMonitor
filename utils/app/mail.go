package app

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
	"strings"
)

func SendMail(mailTo string, subject string, body string) {
	m := gomail.NewMessage()
	//设置发件人
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	//设置发送给多个用户
	mailArrTo := strings.Split(mailTo, ",")
	m.SetHeader("To", mailArrTo...)
	//设置邮件主题
	m.SetHeader("Subject", subject)
	//设置邮件正文
	m.SetBody("text/html", body)
	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	dialer := gomail.NewDialer(os.Getenv("EMAIL_HOST"), port, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS"))
	Error("发送邮件", dialer.DialAndSend(m))
}
