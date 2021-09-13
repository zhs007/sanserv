package world

import (
	"io/ioutil"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type MapConfig struct {
	BasicMapConfig `yaml:",inline"`
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
