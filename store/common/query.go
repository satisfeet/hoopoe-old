package common

import (
	"database/sql"
	"reflect"
	"strconv"
	"strings"

	"github.com/satisfeet/hoopoe/utils"
)

type Query struct {
	err  error
	rows *sql.Rows
}

func (q *Query) All(models interface{}) error {
	defer q.rows.Close()

	if q.err != nil {
		return q.err
	}

	return scanSlice(q.rows, models)
}

func (q *Query) One(model interface{}) error {
	defer q.rows.Close()

	if q.err != nil {
		return q.err
	}

	return scanStruct(q.rows, model)
}

func scanSlice(r *sql.Rows, models interface{}) error {
	c, err := r.Columns()
	if err != nil {
		return err
	}

	for r.Next() {
		m := utils.GetNewType(models)

		if err := scan(r, c, m); err != nil {
			return err
		}

		utils.AppendSlice(models, m)
	}

	return r.Err()
}

func scanStruct(r *sql.Rows, model interface{}) error {
	c, err := r.Columns()
	if err != nil {
		return err
	}

	if r.Next() {
		if err := scan(r, c, model); err != nil {
			return err
		}
	}

	return r.Err()
}

func scan(r *sql.Rows, keys []string, model interface{}) error {
	b := make([]sql.RawBytes, len(keys))
	c := make([]interface{}, len(keys))
	m := make(map[string]sql.RawBytes)

	for i, _ := range b {
		c[i] = &b[i]
	}

	if err := r.Scan(c...); err != nil {
		return err
	}

	for i, k := range keys {
		m[k] = b[i]
	}

	return mapToStruct(m, model)
}

func mapToStruct(m map[string]sql.RawBytes, model interface{}) error {
	t := reflect.Indirect(reflect.ValueOf(model)).Type()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := f.Name

		v := m[strings.ToLower(n)]

		switch f.Type.Kind() {
		case reflect.Slice:
			switch f.Type.Elem().Kind() {
			case reflect.String:
				utils.SetFieldValue(model, n, strings.Split(string(v), ","))
			}
		case reflect.Struct:
			mapToStruct(m, utils.GetFieldPointer(model, n))
		case reflect.String:
			utils.SetFieldValue(model, n, string(v))
		case reflect.Int:
			i, err := strconv.ParseInt(string(v), 0, 32)
			if err != nil {
				return err
			}

			utils.SetFieldValue(model, n, int(i))
		case reflect.Int64:
			i, err := strconv.ParseInt(string(v), 0, 32)
			if err != nil {
				return err
			}

			utils.SetFieldValue(model, n, i)
		case reflect.Float32, reflect.Float64:
			i, err := strconv.ParseFloat(string(v), 64)
			if err != nil {
				return err
			}

			utils.SetFieldValue(model, n, i)
		}

		if scanner, ok := utils.GetFieldPointer(model, n).(sql.Scanner); ok {
			if err := scanner.Scan(v); err != nil {
				return err
			}
		}
	}

	return nil
}
