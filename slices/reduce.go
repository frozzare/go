package slices

import (
	"reflect"
)

// Reduce executes a reducer function on each element of the array, resulting in single output value.
func Reduce(input interface{}, predicate interface{}, accumulator interface{}) (interface{}, error) {
	pv := reflect.ValueOf(predicate)

	each(input, func(c interface{}) interface{} {
		var accumulatorValue reflect.Value
		if accumulator == nil {
			accumulatorValue = reflect.New(reflect.TypeOf(c)).Elem()
		} else {
			accumulatorValue = reflect.ValueOf(accumulator)
		}

		args := []reflect.Value{accumulatorValue, reflect.ValueOf(c)}
		res := pv.Call(args)

		if len(res) > 0 {
			return res[0].Interface()
		}

		return nil
	}, func(current, key, value reflect.Value) bool {
		accumulator = current.Interface()
		return false
	})

	return accumulator, nil
}
