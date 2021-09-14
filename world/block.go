package world

import (
	"github.com/zhs007/sanserv/basic"
)

type Block struct {
	X, Y       int
	BlockType  BlockType
	UserID     basic.UserID
	AllianceID basic.AllianceID
	ResType    basic.ResType
	ResLevel   basic.ResLevel
}

func NewBlock(x, y int, blockType BlockType) *Block {
	return &Block{
		X:         x,
		Y:         y,
		BlockType: blockType,
	}
}
