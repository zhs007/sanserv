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
	ResType           basic.ResType    `yaml:"restype"`
	Levels            []basic.ResLevel `yaml:"levels"`
	LevelPercentageEx []int            `yaml:"levelpercentageex"`
	LevelPercentage   []float32        `yaml:"levelpercentage"`
	Percentage        float32          `yaml:"percentage"`
}

type MapConfig struct {
	BasicMapConfig           `yaml:",inline"`
	IsNoIsland               bool              `yaml:"isnoisland"`
	Mirrors                  int               `yaml:"mirrors"`
	MirrorType               MirrorType        `yaml:"mirrortype"`
	CenterType               CenterType        `yaml:"centertype"`
	Resources                []MapResConfig    `yaml:"resources"`
	TotalResourcesPercentage float32           `yaml:"-"`
	MapBlockType             [][]BlockType     `yaml:"mapblocktype"`
	MapResType               [][]basic.ResType `yaml:"-"`
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
		cl := []BlockType{}

		for x := 0; x < cfg.Width; x++ {
			cl = append(cl, BlockTypeFlat)
		}

		cfg.MapBlockType = append(cfg.MapBlockType, cl)
	}

	return nil
}

// GetBlockState - get a blocktype / restype / reslevel
func (cfg *MapConfig) GetBlockState(x, y int) (BlockType, basic.ResType, basic.ResLevel) {
	return cfg.MapBlockType[y][x], -1, 0
}

func (cfg *MapConfig) genResType(x, y int) basic.ResType {
	cr := rand.Float32()

	for _, v := range cfg.Resources {
		if cr < v.Percentage {
			return v.ResType
		}

		cr -= v.Percentage
	}

	return -1
}

func (cfg *MapConfig) init() error {
	cfg.TotalResourcesPercentage = 0

	for _, v := range cfg.Resources {
		cfg.TotalResourcesPercentage += v.Percentage
	}

	return nil
}
