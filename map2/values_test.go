package map2

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestValues(t *testing.T) {
	v, err := Values(m)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(v.([]int)))

	v, err = Keys(m2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(v.([]interface{})))
}
