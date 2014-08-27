package store

import (
	"encoding/json"
	"time"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

var orderName = "orders"

var orderUnique = []string{
	"number",
}

type Order struct {
	Id       interface{} `bson:"_id"`
	State    OrderState
	Items    []OrderItem
	Pricing  Pricing
	Sequence int         `validate:"required,min=1"`
	Customer interface{} `validate:"required"`
}

type OrderItem struct {
	Product   interface{} `validate:"required"`
	Quantity  int         `validate:"required,min=1"`
	Pricing   Pricing
	Variation Variation
}

type OrderState struct {
	Created time.Time
	Cleared time.Time
	Shipped time.Time
}

func (o Order) Validate() error {
	return validation.Validate(o)
}

func (o Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(o))
}

func NewOrderQuery() *common.Query {
	return common.NewQuery()
}

type OrderStore struct {
	*common.Store
}

func NewOrderStore(s *common.Session) (*OrderStore, error) {
	os := &OrderStore{
		Store: common.NewStore(common.Config{
			Name:   orderName,
			Unique: orderUnique,
		}, s),
	}

	return os, os.Index()
}

func (s *OrderStore) Insert(o *Order) error {
	i, err := s.Sequence().New()

	if err != nil {
		return err
	}

	o.Sequence = i

	o.State.Created = time.Now()

	return s.Store.Insert(o)
}
