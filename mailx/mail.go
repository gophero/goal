package mailx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gophero/goal/logx"
	"github.com/gophero/goal/random"
	"github.com/gophero/goal/redisx"
	"io"
	"net/http"
	"regexp"
	"time"

	"gopkg.in/gomail.v2"
)

type Conf struct {
	Server        string `yaml:"server"`
	SupportDomain string `yaml:"supportDomain"`
}

// Server 邮件服务配置
type Server struct {
	Host     string
	Port     int
	From     string
	UserName string
	Password string
}

func (c Server) CheckValid() error {
	if c.Host == "" || c.Port == 0 || c.UserName == "" || c.Password == "" {
		return errors.New("bad parameters")
	}
	if !ValidMailAddress(c.From) {
		return errors.New("invalid from address: " + c.From)
	}
	return nil
}

// ISender 邮件发送接口
type ISender interface {
	Send(config Server, param Param) error
}

// Param 发送参数
type Param struct {
	To      []string
	Subject string
	Body    string
	Cc      []string
	Bcc     []string
}

func (p Param) CheckValid() error {
	if p.To == nil || len(p.To) == 0 {
		return errors.New("need send to address")
	}
	if p.Subject == "" {
		return errors.New("need subject")
	}
	if p.Body == "" {
		return errors.New("need body")
	}
	// if !ValidMailAddress(p.To) {
	//	return errors.New("invalid to address")
	// }
	return nil
}

func NewSender(log *logx.Logger, conf Conf) *Sender {
	return &Sender{Log: log, conf: conf}
}

// Sender 邮件发送实现
type Sender struct {
	Log  *logx.Logger
	conf Conf
}

func (s *Sender) Send(server Server, param Param) error {
	err := server.CheckValid()
	if err != nil {
		return err
	}
	err = param.CheckValid()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", server.UserName)
	m.SetHeader("To", param.To...)
	m.SetHeader("Cc", param.Cc...)
	m.SetHeader("Bcc", param.Bcc...)
	m.SetHeader("Subject", param.Subject)
	m.SetBody("text/html", param.Body)

	d := gomail.NewDialer(server.Host, server.Port, server.UserName, server.Password)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

var mailAddrRegex, _ = regexp.Compile("^(\\w|\\.|-)+@(\\w|-)+\\.(\\w|-)+(\\w|-|\\.)*$")

func ValidMailAddress(addr string) bool {
	return mailAddrRegex.Match([]byte(addr))
}

type SendReq struct {
	Appid   string         `json:"appid"`
	Code    string         `json:"code"`
	Subject string         `json:"subject"`
	To      []string       `json:"to"`
	Params  map[string]any `json:"params"`
}

func (s *Sender) SendVerifyCode(redis redisx.Client, to string, Type string) {
	verifyCode := random.Numeric(6)
	s.Log.Infof("email: %v, verify code: %v", to, verifyCode)
	// 存储到redis
	redis.Set(context.Background(), to, verifyCode, time.Minute*30)

	req := &SendReq{
		Appid:   "testapp",
		Code:    "email_verify_code_en",
		Subject: "Please Check Your Verify Code",
		To:      []string{to},
		Params:  map[string]any{"code": verifyCode},
	}
	url := s.conf.Server + "/send"
	ct := "application/json"
	body, err := json.Marshal(req)
	if err != nil {
		s.Log.Errorf("error: %v", err)
	}

	resp, err := http.Post(url, ct, bytes.NewReader(body))
	if err != nil {
		s.Log.Errorf("send email error, addr: %s, err: %v", to, err)
		return
	}

	s.Log.Debug("status code: ", resp.StatusCode) // 获取状态码
	s.Log.Debug("status: ", resp.Status)          // 获取状态码对应的文案
	bs, _ := io.ReadAll(resp.Body)                // 读取响应 body, 返回为 []byte
	s.Log.Debug("response: ", string(bs))         // 转成字符串看一下结果
}
