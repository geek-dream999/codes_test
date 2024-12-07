package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

func sendEmail(from string, to []string) {
	e := &email.Email{
		To:      to,
		From:    fmt.Sprintf("测试邮箱 <%s>", from),
		Subject: "测试邮件 2 - 请忽略",
		Text:    []byte("这是一封简单的测试邮件，请忽略。"),
		HTML:    []byte("<h1>这是一封简单的测试邮件，请忽略。</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	//err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "l397608301@gmail.com", "krledtfteffshupv", "smtp.gmail.com"))
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "l397608301@gmail.com", "yzmiplotbzonyduv", "smtp.gmail.com"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("发送成功")
}
