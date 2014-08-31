package store

import (
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
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
	store *common.Store
}

func NewCustomerStore(s *common.Session) *CustomerStore {
	return &CustomerStore{
		store: common.NewStore(s),
	}
}

func (s *CustomerStore) Find(m *[]Customer) error {
	return s.store.Select(`
		SELECT *
		FROM customer_address_city
	`, m)
}

func (s *CustomerStore) FindId(id string, m *Customer) error {
	return s.store.SelectOne(`
		SELECT *
		FROM customer_address_city
		WHERE id=?
	`, m, id)
}
