package world

type MapConfig struct {
	BasicMapConfig `yaml:",inline"`
}

func (cfg *MapConfig) RandFirst() error {
	return nil
}

func (cfg *MapConfig) GenBlockType(x, y int) BlockType {
	if cfg.MapType == MapTypeMainLand {
		return BlockTypeFlat
	}

	return BlockTypeOcean
}
