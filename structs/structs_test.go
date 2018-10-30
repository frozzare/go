package structs

import (
	"testing"

	"github.com/frozzare/go-assert"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Child     *Child
}

type Person2 struct {
	Child Child
}

func TestFieldChild(t *testing.T) {
	s := &Person{
		FirstName: "Fredrik",
		Child: &Child{
			Name: "test",
		},
	}

	f, err := Field(s, "Child")
	assert.Nil(t, err)

	f2, err := f.Field("Name")
	assert.Nil(t, err)

	assert.Equal(t, "test", f2.Value())
}

func TestFieldChildNilPtr(t *testing.T) {
	s := &Person{
		FirstName: "Fredrik",
	}

	f, err := Field(s, "Child")
	assert.Nil(t, err)

	f2, err := f.Field("Name")
	assert.Nil(t, err)

	err = f2.Set("test")
	assert.Nil(t, err)

	assert.Equal(t, "test", f2.Value())
}

func TestFieldChildStruct(t *testing.T) {
	s := &Person2{}

	f, err := Field(s, "Child")
	assert.Nil(t, err)

	f2, err := f.Field("Name")
	assert.Nil(t, err)

	err = f2.Set("test")
	assert.Nil(t, err)

	assert.Equal(t, "test", f2.Value())
}

func TestFieldChildStruct2(t *testing.T) {
	s := &Person2{
		Child: Child{
			Name: "Fredrik",
		},
	}

	f, err := Field(s, "Child")
	assert.Nil(t, err)

	f2, err := f.Field("Name")
	assert.Nil(t, err)

	assert.Equal(t, "Fredrik", f2.Value())

	err = f2.Set("test")
	assert.Nil(t, err)

	assert.Equal(t, "test", f2.Value())
}

func TestField(t *testing.T) {
	s := &Person{
		FirstName: "fredrik",
	}

	f, err := Field(s, "FirstName")
	assert.Nil(t, err)

	assert.Equal(t, "FirstName", f.Name())
	assert.Equal(t, "fredrik", f.Value())
	assert.Equal(t, "first_name", f.Tag("json"))
	assert.Equal(t, "string", f.Kind().String())
	assert.False(t, f.IsZero())

	err = f.Set("elli")
	assert.Nil(t, err)
	assert.Equal(t, "elli", f.Value())
	assert.Equal(t, "elli", s.FirstName)

	f2, err := Field(s, "LastName")
	assert.Nil(t, err)

	assert.True(t, f2.IsZero())
}

func TestFields(t *testing.T) {
	s := &Person{
		FirstName: "fredrik",
	}

	fields, err := Fields(s)
	assert.Nil(t, err)

	assert.Equal(t, "FirstName", fields[0].Name())
}

func TestName(t *testing.T) {
	s := &Person{}
	v, err := Name(s)
	assert.Nil(t, err)
	assert.Equal(t, "Person", v)
}

func TestNames(t *testing.T) {
	s := &Person{}
	v, err := Names(s)

	assert.Nil(t, err)
	assert.Equal(t, []string{"FirstName", "LastName", "Child"}, v)
}
