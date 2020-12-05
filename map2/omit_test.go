package map2

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestOmit(t *testing.T) {
	v, err := Omit([]string{"a", "d"}, map[string]int{"a": 1, "b": 2, "c": 3, "d": 4})

	assert.Nil(t, err)
	assert.Equal(t, map[string]int{"b": 2, "c": 3}, v.(map[string]int))
}
