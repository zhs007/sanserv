package world

type BlockType int

const (
	BlockTypeNone      = 0 // None Type
	BlockTypeFlat      = 1 // flat ground
	BlockTypeMountain  = 2 // mountain
	BlockTypeRiver     = 3 // river
	BlockTypeBridge    = 4 // bridge
	BlockTypePlankRoad = 5 // plank road
	BlockTypeLake      = 6 // lake
	BlockTypeOcean     = 7 // ocean
)

type MapType int

const (
	MapTypeNone     = 0 // None Type
	MapTypeMainLand = 1 // main land
)

type MirrorType int

const (
	MirrorTypeX  = 0 // x/-x
	MirrorTypeY  = 1 // y/-y
	MirrorTypeXY = 2 // x,y/-x,-y
)
