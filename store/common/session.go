package common

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
