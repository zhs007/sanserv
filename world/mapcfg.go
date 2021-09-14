package world

import (
	"io/ioutil"

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
	BasicMapConfig `yaml:",inline"`
	IsNoIsland     bool           `yaml:"isnoisland"`
	Mirrors        int            `yaml:"mirrors"`
	MirrorType     MirrorType     `yaml:"mirrortype"`
	CenterType     CenterType     `yaml:"centertype"`
	Resources      []MapResConfig `yaml:"resources"`
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

	return cfg, nil
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
