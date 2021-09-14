package basic

type UserID int64

type AllianceID int64

type ResType int

const (
	ResTypeMoney = ResType(0)
	ResTypeFood  = ResType(1)
	ResTypeWood  = ResType(2)
	ResTypeStone = ResType(3)
	ResTypeIron  = ResType(4)
)

type ResLevel int
