package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 装饰器模式

type IDraw interface {
	Draw() string
}

// 正方形
type Square struct{}

func (s Square) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{square: square, color: color}
}

func (c ColorSquare) Draw() string {
	return c.square.Draw() + ", color is " + c.color
}

func TestColorSquare(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "red")
	got := csq.Draw()
	assert.Equal(t, "this is a square, color is red", got)
}
