package map2

import (
	"fmt"
	"reflect"
)

func mp(s interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Map {
		return v, fmt.Errorf("%s is not a map", v.Kind())
	}

	return v, nil
}
