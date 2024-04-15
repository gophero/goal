package mailx

import (
	"fmt"
	"testing"
)

func TestValidMailAddress(t *testing.T) {
	type Case struct {
		addr  string
		valid bool
	}
	cases := []Case{
		{"zz871899125@gmail.com", true},
		{"1313831783@qq.com", true},
		{"belonk@126.com", true},
		{"sam.koobyte@gmail.com", true},
		{"t_123@test.com", true},
		{"t-1@test.com", true},
		{"t-1@test.haha.com", true},
		{"t-1@test-abc.haha.com", true},
		{"t-1@test_abc.ha_ha.com", true},
		{"t-1@test_abc.ha-ha.com", true},

		{"t_123com", false},
		{"t_123test.com", false},
		{"t^123@test.com", false},
		{"t*123@test.com", false},
		{"t@123@test.com", false},
	}
	for _, tc := range cases {
		if ValidMailAddress(tc.addr) != tc.valid {
			t.Errorf("Invalid result, mail: %s, expect: %t", tc.addr, tc.valid)
		}
	}
}

func prepareAccount() Server {
	var server Server
	server.Host = "smtp.mail.com"
	// server.Port = 25
	server.Port = 465 // ssl端口
	server.From = "test@mail.com"
	server.UserName = "test@mail.com"
	server.Password = "123"
	return server
}

func prepareParam() Param {
	var param Param
	to := [1]string{"test@mail.com"}
	cc := [1]string{"test1@mail.com"}
	bcc := [1]string{"test2@mail.com"}
	param.To = to[:]
	param.Subject = "test mail"
	param.Cc = cc[:]
	param.Bcc = bcc[:]
	return param
}

func TestSendPlainTextMail(t *testing.T) {
	server := prepareAccount()
	if server.UserName == "" || server.Password == "" {
		t.Skip("skip test because no mail account set.")
		return
	}

	param := prepareParam()
	param.Body = "这是一封测试邮件，请勿回复！<br> This is a test email, please do not reply."

	var sender = &Sender{}
	err := sender.Send(server, param)
	if err != nil {
		t.Errorf("Send email error: %v, expect no error", err)
	}
}

func TestHtmlEmail(t *testing.T) {
	server := prepareAccount()
	if server.UserName == "" || server.Password == "" {
		t.Skip("skip test because no mail account set.")
		return
	}

	param := prepareParam()
	param.Body = "这是一封测试邮件，请勿回复！<br> This is a test email, please do not reply，点击 <a href='https://belonk.com'>这里</a> 查看更多信息."

	var sender = &Sender{}
	err := sender.Send(server, param)
	if err != nil {
		t.Errorf("Send email error: %v, expect no error", err)
	}
}

func TestValidEmail1(t *testing.T) {
	fmt.Println(ValidMailAddress("ukwoma.i.police@gmail.com"))
	fmt.Println(ValidMailAddress("Kejainfrancis@gmail.com"))
}
