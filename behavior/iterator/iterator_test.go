package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 迭代器模式

// 迭代器接口
type Iterator interface {
	HasNext() bool
	Next()
	CurrentItem() interface{} // 获取当前元素，可以使用泛型代替
}

type ArrayInt []int

// 创建迭代器
func (a ArrayInt) Iterator() Iterator {
	return &ArrayIntIterator{
		arrayInt: a,
		index:    0,
	}
}

// 数组迭代
type ArrayIntIterator struct {
	arrayInt ArrayInt
	index    int
}

func (i *ArrayIntIterator) HasNext() bool {
	return i.index < len(i.arrayInt)-1
}

func (i *ArrayIntIterator) Next() {
	i.index++
}

func (i *ArrayIntIterator) CurrentItem() interface{} {
	return i.arrayInt[i.index]
}

func TestIterator(t *testing.T) {
	data := ArrayInt{1, 3, 5, 7, 8}
	iter := data.Iterator()
	i := 0
	for iter.HasNext() {
		assert.Equal(t, data[i], iter.CurrentItem())
		iter.Next()
		i++
	}
}
