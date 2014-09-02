package model

import (
	"database/sql"
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/model/store"
	"github.com/satisfeet/hoopoe/model/utils"
)

type Product struct {
	Id          int64
	Title       string
	Subtitle    string
	Description string
	Categories  []string
	Variations  Variations
	Pricing     Pricing
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

var sqlSelectProduct = `
	SELECT id, title, subtitle, description, price AS retail, variations, categories
	FROM product_variation_category
`

var sqlSelectProductId = `
	SELECT id, title, subtitle, description, price AS retail, variations, categories
	FROM product_variation_category
	WHERE id=?
`

type ProductStore struct {
	store *store.Store
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		store: store.NewStore(db),
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	return s.store.Query(sqlSelectProduct).All(m)
}

func (s *ProductStore) FindId(id interface{}, m *Product) error {
	return s.store.Query(sqlSelectProductId, id).One(m)
}
