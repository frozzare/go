package map2

import (
	"errors"
	"reflect"
)

// Keys return the given map keys or error.
func Keys(m interface{}) (interface{}, error) {
	var slice reflect.Value

	s, err := mp(m)
	if err != nil {
		return nil, err
	}

	for _, k := range s.MapKeys() {
		if !slice.IsValid() {
			typ := reflect.SliceOf(k.Type())
			slice = reflect.MakeSlice(typ, 0, 0)
		}

		slice = reflect.Append(slice, k)
	}

	if slice.IsValid() {
		return slice.Interface(), nil
	}

	return nil, errors.New("No keys found")
}
