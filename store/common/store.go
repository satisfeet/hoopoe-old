package common

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Query(query string, params ...interface{}) *Query {
	rows, err := s.db.Query(query, params...)

	return &Query{
		err:  err,
		rows: rows,
	}
}

func (s *Store) Begin() *Tx {
	tx, err := s.db.Begin()

	return &Tx{
		Tx:  tx,
		err: err,
	}
}
