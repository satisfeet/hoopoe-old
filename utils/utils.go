package utils

import (
	"reflect"
	"strings"
)

// FieldInfo stores all possible information stored on a struct field via tags.
type FieldInfo struct {
	Name   string
	Index  bool
	Unique bool
}

// Tag name to lookup.
const TagName = "store"

// Returns the type name
func GetTypeName(model interface{}) string {
	t := reflect.Indirect(reflect.ValueOf(model)).Type()

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()
	}

	return t.Name()
}

// Returns the interface value of a field.
func GetFieldValue(model interface{}, name string) interface{} {
	v := reflect.Indirect(reflect.ValueOf(model))

	return v.FieldByName(name).Interface()
}

// Sets the interface value of a field.
func SetFieldValue(model interface{}, name string, value interface{}) {
	v := reflect.Indirect(reflect.ValueOf(model))

	if f := v.FieldByName(name); f.CanSet() {
		f.Set(reflect.ValueOf(value))
	}
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

// Extracts recursively all valid store tags from given reflect type.
func getTypeInfo(t reflect.Type) map[string]FieldInfo {
	si := make(map[string]FieldInfo)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := strings.ToLower(f.Name)

		if t := f.Tag.Get(TagName); len(t) > 0 {
			si[n] = FieldInfo{
				Name:   f.Name,
				Index:  strings.Contains(t, "index"),
				Unique: strings.Contains(t, "unique"),
			}
		}

		if f.Type.Kind() == reflect.Struct {
			for k, i := range getTypeInfo(f.Type) {
				si[n+"."+k] = i
			}
		}
	}

	return si
}

// Returns a map of field infos representing tag infos from a slice, array,
// pointer or struct.
func GetStructInfo(model interface{}) map[string]FieldInfo {
	v := reflect.Indirect(reflect.ValueOf(model))

	switch t := v.Type(); t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()

		fallthrough
	default:
		return getTypeInfo(t)
	}
}
