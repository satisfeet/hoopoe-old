package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const Driver = "mysql"

type Session struct {
	database *sql.DB
}

func (s *Session) Dial(url string) error {
	db, err := sql.Open(Driver, url)

	if err != nil {
		return err
	}

	s.database = db

	return nil
}

func (s *Session) Close() error {
	return s.database.Close()
}

func (s *Session) Select(query string, models interface{}) error {
	return s.sqlx().Select(models, query)
}

func (s *Session) SelectOne(query string, model interface{}) error {
	return s.sqlx().Get(model, query)
}

func (s *Session) sqlx() *sqlx.DB {
	return sqlx.NewDb(s.database, Driver)
}
