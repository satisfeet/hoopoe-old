package store

import (
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	Id    int64
	Name  string `validate:"required,min=5"`
	Email string `validate:"required,email"`
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}

type CustomerQuery struct {
}

type CustomerStore struct {
	*common.Store
}

func NewCustomerStore(s *common.Session) *CustomerStore {
	return &CustomerStore{
		Store: common.NewStore(s),
	}
}

func (s *CustomerStore) Find() error {

}
