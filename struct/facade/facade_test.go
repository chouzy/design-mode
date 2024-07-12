package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 门面模式

// 用户接口
type IUser interface {
	Login(phone, code int) (*User, error)
}

// 门面模式
type IUserFacade interface {
	LoginOrRegister(phone, code int) error
}

// 用户
type User struct {
	Name string
}

type UserService struct{}

func (u UserService) Login(phone, code int) (*User, error) {
	return &User{Name: "test login"}, nil
}

func (u UserService) Register(phone, code int) (*User, error) {
	return &User{Name: "test register"}, nil
}

func (u UserService) LoginOrRegister(phone, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return u.Register(phone, code)
}

func TestLogin(t *testing.T) {
	service := UserService{}
	user, err := service.Login(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestRegister(t *testing.T) {
	service := UserService{}
	user, err := service.Register(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}

func TestLoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(13001010101, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
