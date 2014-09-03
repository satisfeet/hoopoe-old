package utils

import (
	"reflect"
	"strings"
)

// Returns new initialized type.
func GetNewType(model interface{}) interface{} {
	t := reflect.Indirect(reflect.ValueOf(model)).Type()

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()
	}

	return reflect.New(t).Interface()
}

func GetStructType(model interface{}) interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))

	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		v = v.Elem()
	}

	return v.Interface()
}

// Returns the interface value pointing to a field.
func GetFieldPointer(model interface{}, name string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v.FieldByName(name).Addr().Interface()
}

func GetNestedFieldPointer(model interface{}, name string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))

	if f := v.FieldByName(name); f.IsValid() {
		return f.Addr().Interface()
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Type().Kind() == reflect.Struct {
			ptr := GetNestedFieldPointer(f.Addr().Interface(), name)

			if ptr != nil {
				return ptr
			}
		}
	}

	return nil
}

// Returns a structs non-zero field values as map with lower case keys.
func GetFieldValues(model interface{}) map[string]interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))
	t := v.Type()

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := strings.ToLower(f.Name)

		// There is a problem when comparing empty values of type
		// []bson.ObjectId() as they do not have a predefined zero value.
		// By checking the array length before we can overgo this however this
		// fix brings a lot of double logic.
		switch v := v.Field(i); v.Kind() {
		case reflect.Slice, reflect.Array:
			if v.Len() > 0 {
				m[n] = v.Interface()
			}
		default:
			if v.Interface() != reflect.Zero(f.Type).Interface() {
				if f.Type.Kind() == reflect.Struct {
					m[n] = GetFieldValues(v.Interface())
				} else {
					m[n] = v.Interface()
				}
			}
		}
	}

	return m
}

// Appends the given model to the given models slice.
func AppendSlice(models interface{}, model interface{}) {
	e := reflect.Indirect(reflect.ValueOf(model))

	s := reflect.ValueOf(models).Elem()
	s.Set(reflect.Append(reflect.Indirect(s), e))
}
