package store

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/satisfeet/hoopoe/utils"
)

var ErrBadScanType = errors.New("bad scan type")

func scanToSlice(r *sql.Rows, models interface{}) error {
	for r.Next() {
		m := utils.GetNewType(models)

		if err := scan(r, m); err != nil {
			return err
		}

		utils.AppendSlice(models, m)
	}

	return r.Err()
}

func scanToStruct(r *sql.Rows, model interface{}) error {
	if r.Next() {
		if err := scan(r, model); err != nil {
			return err
		}
	}

	return r.Err()
}

func scan(r *sql.Rows, model interface{}) error {
	c, err := r.Columns()

	if err != nil {
		return err
	}

	p := make([]interface{}, len(c))

	for i, k := range c {
		k = strings.Title(k)

		p[i] = utils.GetFieldPointer(model, k)
	}

	return r.Scan(p...)
}

func execPrepare(tx *sql.Tx, sql string, params ...interface{}) (int64, error) {
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
