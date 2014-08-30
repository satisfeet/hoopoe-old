package store

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	Id    int64
	Name  string  `validate:"required,min=5"`
	Email *string `validate:"required,email"`
	Address
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}

type CustomerStore struct {
	session *Session
}

func NewCustomerStore(s *Session) *CustomerStore {
	return &CustomerStore{
		session: s,
	}
}

func (s *CustomerStore) Find(m *[]Customer) error {
	return s.sqlx().Select(m, `SELECT * FROM customer_address_city`)
}

func (s *CustomerStore) sqlx() *sqlx.DB {
	return sqlx.NewDb(s.session.database, Driver)
}
