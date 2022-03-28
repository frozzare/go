package pipe

import (
	"strconv"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestMap(t *testing.T) {
	p := Pipe(
		Map(func(item interface{}) int {
			return Value(strconv.Atoi(item.(string))).(int)
		}),
	)

	v, err := p([]string{"1"})

	assert.Nil(t, err)
	assert.Equal(t, v.([]int), []int{1})
}

func TestReduce(t *testing.T) {

}
