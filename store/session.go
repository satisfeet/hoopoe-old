package store

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const Driver = "mysql"

type Session struct {
	database *sql.DB
}

func Open(url string) (*Session, error) {
	db, err := sql.Open(Driver, url)

	if err != nil {
		return nil, err
	}

	return &Session{db}, nil
}

func (s *Session) Close() error {
	return s.database.Close()
}
