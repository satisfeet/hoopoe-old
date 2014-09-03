package httpd

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-validation"
)

type Context struct {
	*context.Context
}

func (c *Context) Error(err error) {
	s := http.StatusInternalServerError

	switch err.(type) {
	case *json.UnmarshalTypeError, validation.Error:
		s = http.StatusBadRequest
	}

	c.Context.Error(err, s)
}

func (c *Context) Respond(value interface{}, code int) {
	v := reflect.Indirect(reflect.ValueOf(value))

	var b interface{}

	switch v.Kind() {
	case reflect.Slice:
		s := make([]map[string]interface{}, v.Len())

		for i := 0; i < v.Len(); i++ {
			s[i] = toMap(v.Index(i))
		}

		b = s
	case reflect.Struct:
		b = toMap(v)
	}

	if err := c.Context.Respond(b, code); err != nil {
		c.Error(err)
	}
}

func toMap(v reflect.Value) map[string]interface{} {
	t := v.Type()

	m := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := strings.ToLower(f.Name)

		switch v := v.Field(i); v.Kind() {
		case reflect.Slice, reflect.Array:
			if v.Len() > 0 {
				m[n] = v.Interface()
			}
		default:
			if v.Interface() != reflect.Zero(f.Type).Interface() {
				if f.Type.Kind() == reflect.Struct {
					m[n] = toMap(v)
				} else {
					m[n] = v.Interface()
				}
			}
		}
	}

	return m
}
