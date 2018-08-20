package template2

import (
	"encoding/json"
	"html/template"
	"reflect"
)

// Isset determine if a map key or struct field is set and valid.
func Isset(data interface{}, name string) bool {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			if key.String() == name && key.IsValid() {
				return v.MapIndex(key).IsValid()
			}
		}
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	return v.FieldByName(name).IsValid()
}

// ToJSON convert the given interface to JSON
// and returns it a template JS type.
func ToJSON(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}
