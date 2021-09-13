package world

type Map struct {
	Width, Height int
	Map           [][]*Block
	Cfg           IMapConfig
}

func NewMap(cfg IMapConfig) *Map {
	m := &Map{
		Width:  cfg.GetWidth(),
		Height: cfg.GetHeight(),
		Cfg:    cfg,
	}

	for y := 0; y < cfg.GetHeight(); y++ {
		ld := []*Block{}

		for x := 0; x < cfg.GetWidth(); x++ {
			cb := NewBlock(x, y, cfg.GenBlockType(x, y))

			ld = append(ld, cb)
		}

		m.Map = append(m.Map, ld)
	}

	return m
}
