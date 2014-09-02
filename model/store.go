package model

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("not found")
var ErrBadScanType = errors.New("bad scan type")

func execPrepare(tx *sql.Tx, sql string, params ...interface{}) error {
	stmt, err := tx.Prepare(sql)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(params...)

	return err
}

func execPrepareId(tx *sql.Tx, sql string, params ...interface{}) (int64, error) {
	stmt, err := tx.Prepare(sql)

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(params...)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func execPrepareAffected(tx *sql.Tx, sql string, params ...interface{}) (int64, error) {
	stmt, err := tx.Prepare(sql)

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(params...)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
