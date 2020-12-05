package slices

import (
	"strconv"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestReduce(t *testing.T) {
	v, err := Reduce([]int{1, 2, 3}, func(p int, v int) int {
		p += v
		return p
	}, 0)

	assert.Nil(t, err)
	assert.Equal(t, 6, v)

	v, err = Reduce([]int{1, 2, 3}, func(p map[string]int, v int) map[string]int {
		p[strconv.Itoa(v)] = v
		return p
	}, map[string]int{})

	res := map[string]int{"1": 1, "2": 2, "3": 3}

	assert.Nil(t, err)
	assert.Equal(t, res, v)
}
