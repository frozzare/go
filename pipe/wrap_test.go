package pipe

import (
	"testing"

	"github.com/frozzare/go-assert"
	"github.com/frozzare/go/slices"
)

var (
	numbers = []int{1, 2, 3}
)

func TestWrap(t *testing.T) {
	p := Pipe(
		Wrap(slices.Map, func(item interface{}) interface{} {
			return item.(int) + 1
		}),
	)

	v, err := p(numbers)

	assert.Nil(t, err)
	assert.Equal(t, []interface{}{2, 3, 4}, v.([]interface{}))
}
