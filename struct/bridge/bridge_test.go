package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 桥接模式

type IMsgSender interface {
	Send(msg string) error
}

// 发送邮件
// 可能还有电话、短信等各种实现
type Email struct {
	emails []string
}

func NewEmail(emails []string) *Email {
	return &Email{emails: emails}
}

func (m *Email) Send(msg string) error {
	return nil
}

// 通知接口
type INotification interface {
	Notify(msg string) error
}

// 错误通知
// 可能还有warning各种级别
type ErrorNotification struct {
	sender IMsgSender
}

func NewError(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

func TestError(t *testing.T) {
	sender := NewEmail([]string{"test@test.com"})
	n := NewError(sender)
	err := n.Notify("test msg")
	assert.Nil(t, err)
}
