// The common package provides utils and interfaces shared between database
// stores.
package common

import (
	"errors"
	"reflect"
	"strings"
)

var (
	// Query specific errors.
	ErrBadQueryId    = errors.New("bad query id")
	ErrBadQueryValue = errors.New("bad query value")

	// Connection specific errors.
	ErrNotConnected   = errors.New("not connected")
	ErrStillConnected = errors.New("still connected")
)

// Query describes a condition container to filter documents.
type Query interface {
	Id(interface{}) error
}

// Stores describes a persistent document storage.
type Store interface {
	// Connection methods.
	Open(string) error
	Close() error

	// Collection methods.
	Drop(string) error

	// Document write methods.
	Insert(string, interface{}) error
	Update(string, Query, interface{}) error
	Remove(string, Query) error

	// Document read methods.
	FindOne(string, Query, interface{}) error
	FindAll(string, Query, interface{}) error
}

// Returns pluralized and lowercased name of type.
func Name(m interface{}) string {
	n := ""

	switch v := value(m); v.Kind() {
	case reflect.Struct:
		n = v.Type().Name()
	case reflect.Array, reflect.Slice:
		n = v.Type().Elem().Name()
	}

	return strings.ToLower(n) + "s"
}

// Returns the reflected value of a pointer or a struct.
func value(m interface{}) reflect.Value {
	if v := reflect.ValueOf(m); v.Kind() == reflect.Ptr {
		return v.Elem()
	} else {
		return v
	}
}
