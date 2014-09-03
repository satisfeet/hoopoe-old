package store

import (
	"database/sql"
	"errors"
)

type Store struct {
	*sql.DB
}

var ErrNotFound = errors.New("not found")
var ErrBadScanType = errors.New("bad scan type")

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) Query(query string, params ...interface{}) *Query {
	var err error
	var keys []string

	rows, err := s.DB.Query(query, params...)

	if err == nil {
		keys, err = rows.Columns()
	}

	return &Query{
		err:  err,
		rows: rows,
		keys: keys,
	}
}

func (s *Store) QueryRows(query string, params ...interface{}) (*sql.Rows, error) {
	return s.DB.Query(query, params...)
}

func (s *Store) Begin() *Tx {
	tx, err := s.DB.Begin()

	return &Tx{
		Tx:  tx,
		err: err,
	}
}
