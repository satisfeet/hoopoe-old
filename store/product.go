package store

import (
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          int64
	Title       string
	Subtitle    string
	Description string
	Price       Pricing
	Categories  Categories
	Variations  Variations
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type ProductQuery struct {
	*common.Query
}

func NewProductQuery() *ProductQuery {
	return &ProductQuery{
		Query: common.NewQuery("product_variation_category"),
	}
}

type ProductStore struct {
	*common.Store
}

func NewProductStore(s *common.Session) *ProductStore {
	return &ProductStore{
		Store: common.NewStore(s),
	}
}
