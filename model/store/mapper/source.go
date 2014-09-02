package mapper

import "database/sql"

type Source []sql.RawBytes

func (s Source) Params() []interface{} {
	p := make([]interface{}, len(s))

	for i, _ := range s {
		p[i] = &s[i]
	}

	return p
}
