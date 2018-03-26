package map2

import (
	"testing"

	"github.com/frozzare/go-assert"
)

var m = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
}

var m2 = map[interface{}]interface{}{
	1:   "2",
	"4": 3,
}

func TestKeys(t *testing.T) {
	v := Keys(m)
	assert.Equal(t, 6, len(v.([]string)))

	v = Keys(m2)
	assert.Equal(t, 2, len(v.([]interface{})))
}

func TestValues(t *testing.T) {
	v := Values(m)
	assert.Equal(t, 6, len(v.([]int)))

	v = Keys(m2)
	assert.Equal(t, 2, len(v.([]interface{})))
}
