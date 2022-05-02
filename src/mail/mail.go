package mail

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

type Sender struct {
	host     string
	port     string
	userName string
	password string
}

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
}

var instanceMap = map[string]*Sender{}

func GetInstance() *Sender {
	return instanceMap["default"]
}

func GetInstance_(name string) *Sender {
	return instanceMap[name]
}

func Init(config Config) error {
	return Init_("default", config)
}

func Init_(name string, config Config) error {
	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("mail sender instance <%s> has already initialized", name)
	}

	if config.Port == 0 {
		config.Port = 587
	}

	sender := &Sender{
		host:     config.Host,
		port:     strconv.Itoa(config.Port),
		userName: config.UserName,
		password: config.Password,
	}

	instanceMap[name] = sender
	return nil
}

func (s *Sender) SendVCode(code string, address string) error {
	auth := smtp.PlainAuth("", s.userName, s.password, s.host)
	e := email.NewEmail()
	e.From = "coreservice <admin@coreservice.io>"
	e.To = []string{address}
	e.Subject = "Verification code"
	e.Text = []byte("Your verification code is [" + code + "], it will expire in 4 hours")

	err := e.Send(s.host+":"+s.port, auth)
	if err != nil {
		return err
	}
	return nil
}
