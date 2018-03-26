package slices

import (
	"testing"

	"github.com/frozzare/go-assert"
)

type Person struct {
	Name string
}

func TestIndexOf(t *testing.T) {
	v := IndexOf([]string{"hello"}, "world")
	assert.Equal(t, -1, v)

	v = IndexOf([]int{1, 2, 3}, 2)
	assert.Equal(t, 1, v)

	v = IndexOf([]Person{Person{Name: "Elli"}, Person{Name: "Fredrik"}}, Person{Name: "Fredrik"})
	assert.Equal(t, 1, v)
}

func TestHas(t *testing.T) {
	v := Has([]string{"hello"}, "world")
	assert.False(t, v)

	v = Has([]int{1, 2, 3}, 2)
	assert.True(t, v)

	v = Has([]Person{Person{Name: "Elli"}, Person{Name: "Fredrik"}}, Person{Name: "Fredrik"})
	assert.True(t, v)
}
