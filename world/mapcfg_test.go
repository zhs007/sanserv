package world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMapConfig(t *testing.T) {
	cfg := &MapConfig{}

	var imc IMapConfig
	imc = cfg
	assert.NotNil(t, imc)

	t.Logf("Test_NewMapConfig OK")
}
