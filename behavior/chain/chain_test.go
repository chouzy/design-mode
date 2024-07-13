package chain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 责任链模式

// 🌰 假设我们现在有个校园论坛，由于社区规章制度、广告、法律法规的原因需要对用户的发言进行敏感词过滤
//    如果被判定为敏感词，那么这篇帖子将会被封禁

// 敏感词过滤
type SensitiveWordFilter interface {
	Filter(content string) bool
}

// 职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

// 添加一个过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

// 执行过滤
func (c *SensitiveWordFilterChain) Filter(context string) bool {
	for _, filter := range c.filters {
		if filter.Filter(context) {
			return true
		}
	}
	return false
}

// 广告
type AdSensitiveWordFilter struct{}

func (f *AdSensitiveWordFilter) Filter(content string) bool {
	// TODO: 实现算法
	return false
}

// 政治敏感
type PoliticalWordFilte struct{}

func (p *PoliticalWordFilte) Filter(content string) bool {
	// TODO: 实现算法
	return true
}

func TestChain(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.Filter("test"))

	chain.AddFilter(&PoliticalWordFilte{})
	assert.Equal(t, true, chain.Filter("test"))
}
