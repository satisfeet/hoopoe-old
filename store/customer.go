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
	return s.sqlx().Select(m, `
		SELECT cu.id, cu.name, cu.email, ad.street, ad.code, ci.name
		FROM customer AS cu
		LEFT JOIN address AS ad ON cu.address_id = ad.id
		LEFT JOIN city AS ci ON ad.city_id = ci.id
	`)
}

func (s *CustomerStore) sqlx() *sqlx.DB {
	return sqlx.NewDb(s.session.database, Driver)
}
