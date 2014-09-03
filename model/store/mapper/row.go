package mapper

import (
	"database/sql"
	"strings"
)

type row struct {
	values []sql.RawBytes
	fields map[string]*sql.RawBytes
}

func newRow(keys []string) *row {
	v := make([]sql.RawBytes, len(keys))
	f := make(map[string]*sql.RawBytes)

	for i, k := range keys {
		f[strings.Title(k)] = &v[i]
	}

	return &row{
		values: v,
		fields: f,
	}
}

func (r *row) Params() []interface{} {
	p := make([]interface{}, len(r.values))

	for i, _ := range r.values {
		p[i] = &r.values[i]
	}

	return p
}
