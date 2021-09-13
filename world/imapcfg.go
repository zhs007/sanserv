package world

type IMapConfig interface {
	// Get Width
	GetWidth() int
	// Get Height
	GetHeight() int
	// Get MapType
	GetMapType() MapType
	// RandFirst - rand at first
	RandFirst() error
	// GenBlockType - generate a blocktype
	GenBlockType(x, y int) BlockType
}

type BasicMapConfig struct {
	Width   int     `yaml:"width"`
	Height  int     `yaml:"height"`
	MapType MapType `yaml:"maptype"`
}

func (cfg BasicMapConfig) GetWidth() int {
	return cfg.Width
}

func (cfg BasicMapConfig) GetHeight() int {
	return cfg.Height
}

func (cfg BasicMapConfig) GetMapType() MapType {
	return cfg.MapType
}
