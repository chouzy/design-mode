package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 懒汉式单例

type Lazy struct{}

var (
	lazy *Lazy
	once = &sync.Once{}
)

func GetLazyInstance() *Lazy {
	if lazy == nil {
		once.Do(func() {
			lazy = &Lazy{}
		})
	}
	return lazy
}

func TestLazy(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}
