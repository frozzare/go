package map2

import (
	"fmt"
	"reflect"
)

func mp(s interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Map {
		return v, fmt.Errorf("%s is not a struct", v.Kind())
	}

	return v, nil
}

// Keys return the given map keys or nil
// if not a map or slice is not valid.
func Keys(m interface{}) interface{} {
	var slice reflect.Value

	s, err := mp(m)
	if err != nil {
		return nil
	}

	for _, k := range s.MapKeys() {
		if !slice.IsValid() {
			typ := reflect.SliceOf(k.Type())
			slice = reflect.MakeSlice(typ, 0, 0)
		}

		slice = reflect.Append(slice, k)
	}

	if slice.IsValid() {
		return slice.Interface()
	}

	return nil
}

// Values return the given map values or nil
// if not a map or slice is not valid.
func Values(m interface{}) interface{} {
	var slice reflect.Value

	s, err := mp(m)
	if err != nil {
		return nil
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
		return slice.Interface()
	}

	return nil
}
