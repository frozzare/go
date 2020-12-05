package map2

import (
	"errors"
	"reflect"
)

// Values return the given map values or error.
func Values(m interface{}) (interface{}, error) {
	var slice reflect.Value

	s, err := mp(m)
	if err != nil {
		return nil, err
	}

	for _, k := range s.MapKeys() {
		v := s.MapIndex(k)

		if !slice.IsValid() {
			typ := reflect.SliceOf(v.Type())
			slice = reflect.MakeSlice(typ, 0, 0)
		}

		slice = reflect.Append(slice, v)
	}

	if slice.IsValid() {
		return slice.Interface(), nil
	}

	return nil, errors.New("No values found")
}
