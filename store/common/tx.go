package common

import "database/sql"

type Tx struct {
	*sql.Tx
	err error
}

type Result struct {
	Id   int64
	Rows int64
}

func (tx *Tx) Exec(query string, params ...interface{}) (Result, error) {
	r := Result{}

	if tx.err != nil {
		return r, tx.err
	}

	stmt, err := tx.Tx.Prepare(query)
	if err != nil {
		tx.Tx.Rollback()

		return r, err
	}

	res, err := stmt.Exec(params...)
	if err != nil {
		tx.Tx.Rollback()

		return r, err
	}

	r.Id, err = res.LastInsertId()
	if err != nil {
		tx.Tx.Rollback()

		return r, err
	}

	r.Rows, err = res.RowsAffected()

	return r, err
}
