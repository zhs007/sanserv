package world

import (
	"github.com/zhs007/sanserv/basic"
)

type Block struct {
	X, Y       int
	BlockType  BlockType
	UserID     basic.UserID
	AllianceID basic.AllianceID
}
