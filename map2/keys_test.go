package map2

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestKeys(t *testing.T) {
	v, err := Keys(m)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(v.([]string)))

	v, err = Keys(m2)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(v.([]interface{})))
}
