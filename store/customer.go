package store

import (
	"database/sql"
	"encoding/json"

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
	db *sql.DB
}

func NewCustomerStore(db *sql.DB) *CustomerStore {
	return &CustomerStore{
		db: db,
	}
}

func (s *CustomerStore) Find(m *[]Customer) error {
	rows, err := s.db.Query(`
		SELECT *
		FROM customer_address_city
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToSlice(rows, m)
}

func (s *CustomerStore) FindId(id interface{}, m *Customer) error {
	rows, err := s.db.Query(`
		SELECT *
		FROM customer_address_city
		WHERE id=?
	`, id)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToStruct(rows, m)
}

func (s *CustomerStore) Search(query string, m *[]Customer) error {
	if len(query) == 0 {
		return s.Find(m)
	}

	query = "'%" + query + "%'"

	rows, err := s.db.Query(`
		SELECT *
		FROM customer_address_city
		WHERE name LIKE ?
		OR email LIKE ?
		OR city LIKE ?
		OR street LIKE ?
	`, query, query, query, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToSlice(rows, m)
}
