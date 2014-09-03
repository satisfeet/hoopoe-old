package store

import (
	"database/sql"

	"github.com/satisfeet/hoopoe/model/store/mapper"
)

type Query struct {
	err  error
	keys []string
	rows *sql.Rows
}

func (q *Query) All(target interface{}) error {
	defer q.rows.Close()

	if q.err != nil {
		return q.err
	}

	m := mapper.NewMapper(target, q.keys)

	for q.rows.Next() {
		if err := scan(q.rows, m); err != nil {
			return err
		}
	}

	return q.rows.Err()
}

func (q *Query) One(target interface{}) error {
	defer q.rows.Close()

	if q.err != nil {
		return q.err
	}

	m := mapper.NewMapper(target, q.keys)

	if q.rows.Next() {
		if err := scan(q.rows, m); err != nil {
			return err
		}
	}

	return q.rows.Err()
}

func scan(r *sql.Rows, m *mapper.Mapper) error {
	src := m.NewSource()

	if err := r.Scan(src.Params()...); err != nil {
		return err
	}

	return m.MapSource(src)
}
