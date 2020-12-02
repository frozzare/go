package file

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/frozzare/go-assert"
)

func TestReadLines(t *testing.T) {
	p := path.Join(t.TempDir(), "readLines")
	d := []byte("hello\ngo\n")

	err := ioutil.WriteFile(p, d, 0644)
	assert.Nil(t, err)

	lines, err := ReadLines(p)
	assert.Nil(t, err)

	assert.Equal(t, len(lines), 2)
	assert.Equal(t, lines[0], "hello")
	assert.Equal(t, lines[1], "go")

	err = os.Remove(p)
	assert.Nil(t, err)
}
