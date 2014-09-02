package store

import (
	"database/sql"

	"github.com/satisfeet/hoopoe/model/store/mapper"
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

	col, err := q.rows.Columns()
	if err != nil {
		return err
	}

	m := mapper.NewMapper(models)
	m.SetColumns(col)

	for q.rows.Next() {
		src := m.NewSource()

		if err := q.rows.Scan(src.Params()...); err != nil {
			return err
		}
		if err := m.MapSource(src); err != nil {
			return err
		}
	}

	return q.rows.Err()
}

func (q *Query) One(model interface{}) error {
	defer q.rows.Close()

	if q.err != nil {
		return q.err
	}

	col, err := q.rows.Columns()
	if err != nil {
		return err
	}

	m := mapper.NewMapper(model)
	m.SetColumns(col)

	if q.rows.Next() {
		src := m.NewSource()

		if err := q.rows.Scan(src.Params()...); err != nil {
			return err
		}
		if err := m.MapSource(src); err != nil {
			return err
		}
	}

	return q.rows.Err()
}
