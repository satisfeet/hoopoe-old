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

type CustomerQuery struct {
	*common.Query
}

func NewCustomerQuery() *CustomerQuery {
	return &CustomerQuery{
		Query: common.NewQuery("customer_address_city"),
	}
}

type CustomerStore struct {
	*common.Store
}

func NewCustomerStore(s *common.Session) *CustomerStore {
	return &CustomerStore{
		Store: common.NewStore(s),
	}
}

func (s *CustomerStore) Find(q *CustomerQuery, m *[]Customer) error {
	return s.Store.Find(q, m)
}

func (s *CustomerStore) FindOne(q *CustomerQuery, m *Customer) error {
	return s.Store.FindOne(q, m)
}
