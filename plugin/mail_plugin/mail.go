package mail_plugin

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/jordan-wright/email"
)

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
}

type Sender struct {
	host     string
	port     int
	userName string
	password string
}

var instanceMap = map[string]*Sender{}

func GetInstance() *Sender {
	return instanceMap["default"]
}

func GetInstance_(name string) *Sender {
	return instanceMap[name]
}

func Init(config *Config) error {
	return Init_("default", config)
}

func Init_(name string, config *Config) error {

	if name == "" {
		name = "default"
	}

	_, exist := instanceMap[name]
	if exist {
		return fmt.Errorf("email instance <%s> has already been initialized", name)
	}

	if config.Port == 0 {
		config.Port = 587
	}

	sender := &Sender{
		host:     config.Host,
		port:     config.Port,
		userName: config.UserName,
		password: config.Password,
	}

	instanceMap[name] = sender
	return nil
}

func (s *Sender) Send(from_text string, to_address string, subject string, body string) error {

	e := email.NewEmail()
	e.From = from_text
	e.To = []string{to_address}
	e.Subject = subject
	e.Text = []byte(body)

	auth := smtp.PlainAuth("", s.userName, s.password, s.host)
	err := e.Send(s.host+":"+strconv.Itoa(s.port), auth)
	if err != nil {
		return err
	}
	return nil
}
