package composite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 组合模式

// 组织接口，实现统计人数的功能
type IOrganization interface {
	Count() int
}

// 员工
type Employee struct {
	Name string
}

func (e *Employee) Count() int {
	return 1
}

// 部门
type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d *Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}
	return root
}

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	assert.Equal(t, 20, got)
}
