package flyweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 享元模式

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "车",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
	// ... 其他棋子
}

// 棋子享元
type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

func NewChessPieceUnit(id int) *ChessPieceUnit {
	return units[id]
}

// 棋子
type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

// 棋局
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

// 初始化棋盘
func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: NewChessPieceUnit(id),
			X:    0,
			Y:    0,
		}
	}
	return board
}

// 移动棋子
func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}

func TestChess(t *testing.T) {
	b1 := NewChessBoard()
	b1.Move(1, 1, 2)
	b2 := NewChessBoard()
	b2.Move(2, 2, 3)

	assert.Equal(t, b1.chessPieces[1].Unit, b2.chessPieces[1].Unit)
	assert.Equal(t, b1.chessPieces[2].Unit, b2.chessPieces[2].Unit)
}
