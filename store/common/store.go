package common

import (
	"errors"
	"reflect"
	"strings"
)

var (
	// Query errors.
	ErrBadQueryId    = errors.New("bad query id")
	ErrBadQueryValue = errors.New("bad query value")
)

type Id interface {
	String() string
}

type Model interface {
	Validate() error
}

type Query interface {
	Id(interface{}) error
}

type Store interface {
	Open(string) error
	Close() error

	Insert(Model) error
	Update(Model) error
	Remove(Model) error

	Find(Query, interface{}) error
	FindOne(Query, Model) error
}

func GetId(m Model) interface{} {
	v := value(m).FieldByName("Id")
	if v, ok := v.Interface().(interface{}); ok {
		return v
	}

	return nil
}

func GetName(m interface{}) string {
	n := ""

	switch v := value(m); v.Kind() {
	case reflect.Struct:
		n = v.Type().Name()
	case reflect.Array, reflect.Slice:
		n = v.Type().Elem().Name()
	}

	return strings.ToLower(n)
}

func value(m interface{}) reflect.Value {
	if v := reflect.ValueOf(m); v.Kind() == reflect.Ptr {
		return v.Elem()
	} else {
		return v
	}
}
