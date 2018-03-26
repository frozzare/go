package sliceutil

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestMap(t *testing.T) {
	v := Map([]int{1, 2, 3}, func(v int) int {
		return v + 1
	})

	assert.Equal(t, []int{2, 3, 4}, v.([]int))

	v2 := Map([]int{1, 2, 3}, func(v, i int) int {
		return v + i
	})

	assert.Equal(t, []int{1, 3, 5}, v2.([]int))
}
