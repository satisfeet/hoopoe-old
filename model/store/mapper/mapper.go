package mapper

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"

	"github.com/satisfeet/hoopoe/model/utils"
)

var KeyFunc = strings.ToLower

type Mapper struct {
	model   interface{}
	models  interface{}
	columns []string
	mappers map[string]MapperFunc
	keyFunc func(string) string
}

var ErrBadSrc = errors.New("bad source type")

func NewMapper(target interface{}) *Mapper {
	m := &Mapper{keyFunc: KeyFunc}
	m.mappers = make(map[string]MapperFunc)
	mappersFrom(reflect.Indirect(reflect.ValueOf(target)).Type(), m.mappers)

	v := reflect.Indirect(reflect.ValueOf(target))

	if v.Type().Kind() == reflect.Slice {
		m.models = target
	} else {
		m.model = target
	}

	return m
}

func (m *Mapper) SetColumns(c []string) {
	m.columns = c
}

func (m *Mapper) SetKeyFunc(fn func(string) string) {
	m.keyFunc = fn
}

func (m *Mapper) NewSource() Source {
	return make(Source, len(m.columns))
}

func (m *Mapper) MapSource(r Source) error {
	var s interface{}

	if m.model == nil {
		s = reflect.New(utils.MustSliceTypeOf(m.models)).Interface()
	} else {
		s = m.model
	}

	for i := 0; i < len(r); i++ {
		k, v := m.columns[i], r[i]

		ptr := utils.MustFieldByName(s, strings.Title(k)).Addr().Interface()

		if err := m.mappers[strings.Title(k)](string(v), ptr); err != nil {
			return errors.New(k + ": " + err.Error())
		}
	}

	if m.model == nil {
		e := reflect.Indirect(reflect.ValueOf(s))

		s := reflect.ValueOf(m.models).Elem()
		s.Set(reflect.Append(reflect.Indirect(s), e))
	}

	return nil
}

func mappersFrom(t reflect.Type, m map[string]MapperFunc) {
	if t.Kind() == reflect.Slice {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := f.Name

		if reflect.PtrTo(f.Type).Implements(reflect.TypeOf((*sql.Scanner)(nil)).Elem()) {
			m[n] = MapScannerFunc

			continue
		}

		switch f.Type.Kind() {
		case reflect.Slice:
			m[n] = MapSliceFunc
		case reflect.String:
			m[n] = MapStringFunc
		case reflect.Struct:
			mappersFrom(f.Type, m)
		case reflect.Int, reflect.Int64:
			m[n] = MapIntFunc
		case reflect.Float32, reflect.Float64:
			m[n] = MapFloatFunc
		}
	}
}
