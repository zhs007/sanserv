package world

type BlockType int

const (
	BlockTypeNone      = BlockType(0) // None Type
	BlockTypeFlat      = BlockType(1) // flat ground
	BlockTypeMountain  = BlockType(2) // mountain
	BlockTypeRiver     = BlockType(3) // river
	BlockTypeBridge    = BlockType(4) // bridge
	BlockTypePlankRoad = BlockType(5) // plank road
	BlockTypeLake      = BlockType(6) // lake
	BlockTypeOcean     = BlockType(7) // ocean
)

type MapType int

const (
	MapTypeNone     = MapType(0) // None Type
	MapTypeMainLand = MapType(1) // main land
)

type MirrorType int

const (
	MirrorTypeX  = MirrorType(0) // x/-x
	MirrorTypeY  = MirrorType(1) // y/-y
	MirrorTypeXY = MirrorType(2) // x,y/-x,-y
)

type CenterType int

const (
	CenterTypeNone     = CenterType(0) // none
	CenterTypeMountain = CenterType(1) // mountain
	CenterTypeLake     = CenterType(2) // lake
)
