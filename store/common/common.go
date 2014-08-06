package common

import (
	"reflect"
	"strings"
)

// Tag name to lookup.
const TagName = "store"

// FieldInfo stores all possible information stored on a struct field via tags.
type FieldInfo struct {
	Name   string
	Index  bool
	Unique bool
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
