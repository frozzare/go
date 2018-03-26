package sliceutil

import (
	"reflect"
)

func filter(input, predicate interface{}, compare bool) interface{} {
	var slice reflect.Value

	each(input, predicate, func(current, key, value reflect.Value) bool {
		if !slice.IsValid() {
			typ := reflect.SliceOf(value.Type())
			slice = reflect.MakeSlice(typ, 0, 0)
		}

		if current.Bool() == compare {
			slice = reflect.Append(slice, value)
		}

		return false
	})

	if slice.IsValid() {
		return slice.Interface()
	}

	return nil
}

// Filter returns a new slice containing all values
// in the slice that satisfy the predicate function.
func Filter(input, predicate interface{}) interface{} {
	return filter(input, predicate, true)
}
