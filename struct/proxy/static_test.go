package proxy

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// 静态代理

type IUser interface {
	Login(username, password string) error
}

// 用户
// @proxy IUser // ！！！使用动态代理时需要加这一行
type User struct{}

// 用户登录
func (u *User) Login(username, password string) error {
	return nil
}

// 代理类
type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

// 和User实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	// 原有业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after 这里可能会有一些收尾逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}

func TestLogin(t *testing.T) {
	proxy := NewUserProxy(&User{})
	err := proxy.Login("test", "password")
	require.Nil(t, err)
}
