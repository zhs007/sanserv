package world

import (
	"io/ioutil"
	"math/rand"

	"github.com/zhs007/goutils"
	"github.com/zhs007/sanserv/basic"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type MapResConfig struct {
	ResType            basic.ResType    `yaml:"restype"`
	Levels             []basic.ResLevel `yaml:"levels"`
	LevelPercentage    []int            `yaml:"levelpercentage"`
	MaxLevelPercentage int              `yaml:"-"`
	Percentage         float32          `yaml:"percentage"`
}

type MapConfig struct {
	BasicMapConfig `yaml:",inline"`
	IsNoIsland     bool               `yaml:"isnoisland"`
	Mirrors        int                `yaml:"mirrors"`
	MirrorType     MirrorType         `yaml:"mirrortype"`
	CenterType     CenterType         `yaml:"centertype"`
	Resources      []MapResConfig     `yaml:"resources"`
	MapBlockType   [][]BlockType      `yaml:"mapblocktype"`
	MapResType     [][]basic.ResType  `yaml:"-"`
	MapResLevel    [][]basic.ResLevel `yaml:"-"`
}

// LoadMapConfig - load from yaml
func LoadMapConfig(fn string) (*MapConfig, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		goutils.Error("LoadMapConfig:ReadFile",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	cfg := &MapConfig{}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		goutils.Error("LoadMapConfig:Unmarshal",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	err = cfg.init()
	if err != nil {
		goutils.Error("LoadMapConfig:init",
			zap.String("fn", fn),
			zap.Error(err))

		return nil, err
	}

	return cfg, nil
}

func (cfg *MapConfig) RandFirst() error {
	for y := 0; y < cfg.Height; y++ {
		cbl := []BlockType{}
		crtl := []basic.ResType{}
		crll := []basic.ResLevel{}

		for x := 0; x < cfg.Width; x++ {
			cbl = append(cbl, BlockTypeFlat)

			rt, rl := cfg.genRes(x, y)

			crtl = append(crtl, rt)
			crll = append(crll, rl)
		}

		cfg.MapBlockType = append(cfg.MapBlockType, cbl)
		cfg.MapResType = append(cfg.MapResType, crtl)
		cfg.MapResLevel = append(cfg.MapResLevel, crll)
	}

	return nil
}

// GetBlockState - get a blocktype / restype / reslevel
func (cfg *MapConfig) GetBlockState(x, y int) (BlockType, basic.ResType, basic.ResLevel) {
	return cfg.MapBlockType[y][x], cfg.MapResType[y][x], cfg.MapResLevel[y][x]
}

func (cfg *MapConfig) genRes(x, y int) (basic.ResType, basic.ResLevel) {
	cr := rand.Float32()

	for _, v := range cfg.Resources {
		if cr < v.Percentage {
			cr1 := rand.Int() % v.MaxLevelPercentage
			for k1, v1 := range v.LevelPercentage {
				if cr1 < v1 {
					return v.ResType, v.Levels[k1]
				}
			}

			return v.ResType, v.Levels[len(v.Levels)-1]
		}

		cr -= v.Percentage
	}

	return -1, 0
}

func (cfg *MapConfig) init() error {
	for k, v := range cfg.Resources {
		cfg.Resources[k].MaxLevelPercentage = 0

		for _, v1 := range v.LevelPercentage {
			cfg.Resources[k].MaxLevelPercentage += v1
		}
	}

	return nil
}
