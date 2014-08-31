package common

import (
	"database/sql"
	"strings"

	"github.com/satisfeet/hoopoe/utils"
)

type Store struct {
	session *Session
}

func NewStore(s *Session) *Store {
	return &Store{
		session: s,
	}
}

func (s *Store) Select(query string, models interface{}, params ...interface{}) error {
	rows, err := s.db().Query(query, params...)

	if err != nil {
		return err
	}

	defer rows.Close()

	k, err := rows.Columns()

	if err != nil {
		return err
	}

	for rows.Next() {
		m := utils.GetNewType(models)

		if err := rows.Scan(toScan(k, m)...); err != nil {
			return err
		}

		utils.AppendSlice(models, m)
	}

	return rows.Err()
}

func (s *Store) SelectOne(query string, model interface{}, params ...interface{}) error {
	rows, err := s.db().Query(query, params...)

	if err != nil {
		return err
	}

	defer rows.Close()

	k, err := rows.Columns()

	if err != nil {
		return err
	}

	if rows.Next() {
		if err := rows.Scan(toScan(k, model)...); err != nil {
			return err
		}
	}

	return rows.Err()
}

func (s *Store) db() *sql.DB {
	return s.session.database
}

func toScan(keys []string, model interface{}) []interface{} {
	p := make([]interface{}, len(keys))

	for i, k := range keys {
		k = strings.Title(k)

		p[i] = utils.GetFieldPointer(model, k)
	}

	return p
}
