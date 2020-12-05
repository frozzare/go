package pipe

import (
	"fmt"
	"testing"

	"github.com/frozzare/go-assert"
)

func addOne(a int) int {
	return a + 1
}

func add(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func addB(a int, b int) int {
	return a + b
}

func multiTen(a int) int {
	return a * 10
}

func toString(a int) string {
	return fmt.Sprintf("%d", a)
}

func TestPipeSimple(t *testing.T) {
	a := Pipe(addOne, multiTen, toString)
	v, err := a(1)
	assert.Nil(t, err)
	assert.Equal(t, "20", v)
}

func TestPipeAddFunction(t *testing.T) {
	a := Pipe(add(1), multiTen, toString)
	v, err := a(1)
	assert.Nil(t, err)
	assert.Equal(t, "20", v)
}

func TestPipeMultipleArguments(t *testing.T) {
	a := Pipe(addB, Tap(fmt.Println), toString)
	v, err := a(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, "3", v)
}

func TestTap(t *testing.T) {
	v := Tap(fmt.Println, 100, "message")

	assert.Nil(t, Error(v))
	assert.Equal(t, 100, v)
}
