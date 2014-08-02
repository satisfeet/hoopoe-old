package common

import (
	"reflect"
	"strings"
)

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

// Tag name to lookup.
const TagName = "store"

// FieldInfo stores all possible information stored on a struct field via tags.
type FieldInfo struct {
	Name   string
	Index  bool
	Unique bool
	Search bool
}

// StructInfo stores all field information for one struct.
type StructInfo struct {
	FieldMap   map[string]FieldInfo
	FieldArray []FieldInfo
}

// Adds a field info to struct.
func (si *StructInfo) addField(fi FieldInfo) {
	if si.FieldMap == nil {
		si.FieldMap = make(map[string]FieldInfo)
	}
	if si.FieldArray == nil {
		si.FieldArray = make([]FieldInfo, 0)
	}

	si.FieldMap[fi.Name] = fi
	si.FieldArray = append(si.FieldArray, fi)
}

// Parsed the tags of the given model and returns a struct info.
func GetStructInfo(model interface{}) *StructInfo {
	si := new(StructInfo)

	t := reflect.Indirect(reflect.ValueOf(model)).Type()

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()

		fallthrough
	case reflect.Ptr:
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if t := f.Tag.Get(TagName); len(t) > 0 {
			si.addField(FieldInfo{
				Name:   strings.ToLower(f.Name),
				Index:  strings.Contains(t, "index"),
				Unique: strings.Contains(t, "unique"),
				Search: strings.Contains(t, "search"),
			})
		}
	}

	return si
}
