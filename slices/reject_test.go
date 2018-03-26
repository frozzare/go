package slices

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestReject(t *testing.T) {
	v := Reject([]int{1, 2, 3}, func(v int) bool {
		return v%2 == 0
	})

	assert.Equal(t, []int{1, 3}, v.([]int))
}
