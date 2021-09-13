package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	t.Logf("Test_LoadMapConfig OK")
}
