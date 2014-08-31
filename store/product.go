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

type ProductStore struct {
	store *common.Store
}

func NewProductStore(s *common.Session) *ProductStore {
	return &ProductStore{
		store: common.NewStore(s),
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	return s.store.Select(`
		SELECT *
		FROM product_variation_category
	`, m)
}

func (s *ProductStore) FindId(id string, m *Product) error {
	return s.store.Select(`
		SELECT *
		FROM product_variation_category
		WHERE id=?
	`, m, id)
}
