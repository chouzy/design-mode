package mediator

import (
	"fmt"
	"reflect"
	"testing"
)

// 中介模式

// 假设我们现在有一个较为复杂的对话框，里面包括，登录组件，注册组件，以及选择框
// 当选择框选择“登录”时，展示登录相关组件
// 当选择框选择“注册”时，展示注册相关组件

// 假设这是一个输入框
type Input string

func (i Input) String() string {
	return string(i)
}

// 假设这是一个选择框
type Selection string

// 当前选中的对象
func (s Selection) Selected() string {
	return string(s)
}

// 假设这表示一个按钮
type Button struct {
	onClick func()
}

// 添加点击事件回调
func (b *Button) SetOnClick(f func()) {
	b.onClick = f
}

// 中介模式接口
type IMediator interface {
	HandleEvent(component interface{})
}

// 对话框组件
type Dialog struct {
	LoginButton         *Button
	RegButton           *Button
	Selection           *Selection
	UsernameInput       *Input
	PasswordInput       *Input
	RepeatPasswordInput *Input
}

func (d *Dialog) HandleEvent(component interface{}) {
	switch {
	case reflect.DeepEqual(component, d.Selection):
		if d.Selection.Selected() == "登录" {
			fmt.Println("select login")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
		} else if d.Selection.Selected() == "注册" {
			fmt.Println("select register")
			fmt.Printf("show: %s\n", d.UsernameInput)
			fmt.Printf("show: %s\n", d.PasswordInput)
			fmt.Printf("show: %s\n", d.RepeatPasswordInput)
		}
		// others, 如果点击了登录按钮，注册按钮
	}
}

func TestMediator(t *testing.T) {
	usernameInput := Input("username input")
	passwordInput := Input("password input")
	repeatPwdInput := Input("repeat password input")

	selection := Selection("登录")
	d := &Dialog{
		Selection:           &selection,
		UsernameInput:       &usernameInput,
		PasswordInput:       &passwordInput,
		RepeatPasswordInput: &repeatPwdInput,
	}
	d.HandleEvent(&selection)

	regSelection := Selection("注册")
	d.Selection = &regSelection
	d.HandleEvent(&regSelection)
}
