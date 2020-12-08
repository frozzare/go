package pipe

import (
	"testing"

	"github.com/frozzare/go-assert"
)

var (
	mp = map[string]interface{}{
		"name": "pipe",
	}
)

func TestPop(t *testing.T) {
	p := Pipe(Prop("name"))
	v, err := p(mp)
	assert.Nil(t, err)
	assert.Equal(t, "pipe", v)
}
