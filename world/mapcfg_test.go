package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhs007/goutils"
	"github.com/zhs007/sanserv/basic"
)

func Test_NewMapConfig(t *testing.T) {
	cfg := &MapConfig{}

	var imc IMapConfig = cfg
	assert.NotNil(t, imc)

	t.Logf("Test_NewMapConfig OK")
}

func Test_LoadMapConfig(t *testing.T) {
	cfg, err := LoadMapConfig("../cfg/mapconfig.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, cfg.Width, 128)
	assert.Equal(t, cfg.Height, 128)
	assert.Equal(t, cfg.MapType, MapType(1))
	assert.Equal(t, cfg.IsNoIsland, true)
	assert.Equal(t, cfg.Mirrors, 2)
	assert.Equal(t, cfg.MirrorType, MirrorType(2))
	assert.Equal(t, cfg.CenterType, CenterType(0))
	assert.Equal(t, len(cfg.Resources), 4)
	assert.Equal(t, cfg.Resources[0].ResType, basic.ResTypeFood)
	assert.Equal(t, len(cfg.Resources[0].Levels), 10)
	assert.Equal(t, len(cfg.Resources[0].LevelPercentageEx), 10)
	assert.Equal(t, goutils.IsFloatEquals(float64(cfg.Resources[0].Percentage), 0.05), true)

	t.Logf("Test_LoadMapConfig OK")
}
