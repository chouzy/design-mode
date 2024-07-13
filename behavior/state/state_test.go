package state

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 状态模式

// 这是一个工作流的例子，在企业内部或者是学校我们经常会看到很多审批流程
// 假设我们有一个报销的流程: 员工提交报销申请 -> 直属部门领导审批 -> 财务审批 -> 结束
// 在这个审批流中，处在不同的环节就是不同的状态
// 而流程的审批、驳回就是不同的事件

// 状态
type IState interface {
	Approval(m *Machine) // 审批通过
	Reject(m *Machine)   // 驳回
	GetName() string     // 获取状态名称
}

// 状态机
type Machine struct {
	state IState
}

// 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

// 获取当前状态
func (m *Machine) GetSteteName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// 直属领导审批
type leader struct{}

func (l leader) Approval(m *Machine) {
	fmt.Println("leader pass")
	m.SetState(GetFinanceApproveState())
}

func (l leader) GetName() string {
	return "leader"
}

func (l leader) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leader{}
}

// 财务
type finance struct{}

func (f finance) Approval(m *Machine) {
	fmt.Println("finance pass")
}

func (f finance) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

func (f finance) GetName() string {
	return "finance"
}

func GetFinanceApproveState() IState {
	return &finance{}
}

func TestMachine(t *testing.T) {
	m := &Machine{state: GetLeaderApproveState()}
	assert.Equal(t, "leader", m.GetSteteName())
	m.Approval()
	assert.Equal(t, "finance", m.GetSteteName())
	m.Reject()
	assert.Equal(t, "leader", m.GetSteteName())
	m.Approval()
	assert.Equal(t, "finance", m.GetSteteName())
	m.Approval()
}
