package store

import (
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

var customerName = "customers"

var customerIndex = []string{
	"company",
	"address.street",
	"address.city",
}

var customerUnique = []string{
	"email",
	"name",
}

type Customer struct {
	Id      interface{} `bson:"_id"`
	Name    string      `validate:"required,min=5"`
	Email   string      `validate:"required,email"`
	Company string      `validate:"min=5,max=40"`
	Address Address     `validate:"required"`
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
		Query: common.NewQuery(),
	}
}

// Applies a pseudo full text search on all fields on the index.
func (q *CustomerQuery) Search(query string) {
	if len(query) == 0 {
		return
	}

	for _, f := range append(customerIndex, customerUnique...) {
		o := common.NewQuery()
		o.RegEx(f, query)

		q.Or(o)
	}
}

type CustomerStore struct {
	*common.Store
}

func NewCustomerStore(s *common.Session) (*CustomerStore, error) {
	cs := &CustomerStore{
		Store: common.NewStore(common.Config{
			Name:   customerName,
			Index:  customerIndex,
			Unique: customerUnique,
		}, s),
	}

	return cs, cs.Index()
}
