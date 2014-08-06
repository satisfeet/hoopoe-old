package utils

import (
	"reflect"
	"strings"
)

// Returns a structs non-zero field values as map with lower case keys.
func GetFieldValues(model interface{}) map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))
	t := v.Type()

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := strings.ToLower(f.Name)

		val := v.Field(i).Interface()

		if val != reflect.Zero(f.Type).Interface() {
			if f.Type.Kind() == reflect.Struct {
				m[n] = GetFieldValues(val)
			} else {
				m[n] = val
			}
		}
	}

	return m
}
