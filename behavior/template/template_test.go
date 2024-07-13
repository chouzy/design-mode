package template

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 模板模式

// ISMS
type ISMS interface {
	send(conten string, phone int) error
}

// SMS 短信发送基类
type sms struct {
	ISMS
}

// 校验短信字数
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content id too long")
	}
	return nil
}

// 发送短信
func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	// 调用子类方法发送短信
	return s.send(content, phone)
}

// 电信通道
type TelecomSms struct {
	*sms
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (t *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}

func TestSms(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 123999)
	assert.NoError(t, err)
}
