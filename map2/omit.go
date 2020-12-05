package map2

import (
	"errors"
	"reflect"

	"github.com/frozzare/go/slices"
)

// Omit returns a partial copy of an map omitting the keys specified.
func Omit(names []string, obj interface{}) (interface{}, error) {
	s, err := mp(obj)
	if err != nil {
		return nil, err
	}

	next := reflect.MakeMapWithSize(reflect.TypeOf(obj), 0)

	for _, k := range s.MapKeys() {
		if !slices.Contains(names, k.Interface()) {
			next.SetMapIndex(k, s.MapIndex(k))
		}
	}

	if next.IsValid() {
		return next.Interface(), nil
	}

	return nil, errors.New("Not a valid slice")
}
