package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 饿汉式单例
type Eager struct{}

var eager *Eager

func init() {
	eager = &Eager{}
}

func GetEagerInstance() *Eager {
	return eager
}

func TestEager(t *testing.T) {
	assert.Equal(t, GetEagerInstance(), GetEagerInstance())
}
