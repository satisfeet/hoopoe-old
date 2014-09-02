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
	rows, err := s.DB.Query(query, params...)

	return &Query{
		err:  err,
		rows: rows,
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
