package utils

import (
	"errors"
	"reflect"
)

var (
	ErrInvalidField  = errors.New("invalid field")
	ErrInvalidSlice  = errors.New("invalid slice")
	ErrInvalidStruct = errors.New("invalid struct")
)

func SliceTypeOf(source interface{}) (reflect.Type, error) {
	t := reflect.Indirect(reflect.ValueOf(source)).Type()

	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()
	default:
		return nil, ErrInvalidSlice
	}

	return t, nil
}

func FieldByName(source interface{}, name string) (reflect.Value, error) {
	v := reflect.Indirect(reflect.ValueOf(source))

	if v.Kind() != reflect.Struct {
		return v, ErrInvalidStruct
	}

	if f := v.FieldByName(name); f.IsValid() {
		return f, nil
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		if f.Type().Kind() == reflect.Struct {
			v, err := FieldByName(f.Addr().Interface(), name)

			if err == nil {
				return v, nil
			}
		}
	}

	return v, ErrInvalidField
}

func MustSliceTypeOf(source interface{}) reflect.Type {
	t, err := SliceTypeOf(source)

	if err != nil {
		panic(err)
	}

	return t
}

func MustFieldByName(source interface{}, name string) reflect.Value {
	v, err := FieldByName(source, name)

	if err != nil {
		panic(err)
	}

	return v
}

func AppendSlice(slice interface{}, elem interface{}) {
	e := reflect.Indirect(reflect.ValueOf(elem))

	s := reflect.ValueOf(slice).Elem()
	s.Set(reflect.Append(reflect.Indirect(s), e))
}
